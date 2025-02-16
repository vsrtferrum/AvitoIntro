// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/vsrtferrum/AvitoIntro/internal/auth"
	"github.com/vsrtferrum/AvitoIntro/internal/cache"
	"github.com/vsrtferrum/AvitoIntro/internal/database"
	internalError "github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/logger"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/module"
	readconfig "github.com/vsrtferrum/AvitoIntro/internal/read_config"
	"github.com/vsrtferrum/AvitoIntro/internal/workers"
	"github.com/vsrtferrum/AvitoIntro/models"
	"github.com/vsrtferrum/AvitoIntro/restapi/operations"
)

var (
	pathToCfg = "cfg/cfg.json"
)

//go:generate swagger generate server --target ../../AvitoIntro --name Avitoapi --spec ../docs/api.json --principal interface{}

func configureFlags(api *operations.AvitoapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AvitoapiAPI) http.Handler {
	// configure the api here
	cacheCfg, dbCfg, logsCfg, authTTL, workersCfg, err := readconfig.ReadConfig(pathToCfg)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	logger, err := logger.NewLogger(logsCfg)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	cache, err := cache.NewCahce(cacheCfg)
	if err != nil {
		logger.WriteError(err)
		return nil
	}

	db := database.NewDatabase(dbCfg, logger)
	err = db.Connect()
	if err != nil {
		logger.WriteError(err)
		return nil
	}

	authLayer := auth.NewAuth(authTTL)
	module := module.NewModule(&cache, &db, logger, &authLayer)
	workers := workers.NewConcurrentModule(&module, workersCfg.WorkersCount, workersCfg.WorkersQueueLen)

	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuthAuth == nil {
		api.BearerAuthAuth = func(token string) (interface{}, error) {
			ok := workers.Identify(token)
			if !ok {
				operations.NewPostAPIAuthInternalServerError()
				return nil, internalError.ErrIdentifyUser
			}
			return ok, nil
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.GetAPIBuyItemHandler == nil {
		api.GetAPIBuyItemHandler = operations.GetAPIBuyItemHandlerFunc(func(params operations.GetAPIBuyItemParams, principal interface{}) middleware.Responder {
			token, err := auth.ExtractTokenFromHeader(params.HTTPRequest)
			if err != nil {
				return operations.NewGetAPIBuyItemUnauthorized()
			}
			err = workers.Buy(token, params.Item)
			if err != nil {
				return operations.NewGetAPIBuyItemInternalServerError()
			}
			return operations.NewGetAPIBuyItemOK()
		})
	}
	if api.GetAPIInfoHandler == nil {
		api.GetAPIInfoHandler = operations.GetAPIInfoHandlerFunc(func(params operations.GetAPIInfoParams, principal interface{}) middleware.Responder {
			token, err := auth.ExtractTokenFromHeader(params.HTTPRequest)
			if err != nil {
				return operations.NewGetAPIInfoUnauthorized()
			}

			res, err := workers.GetInfo(token)
			if err != nil {
				return operations.NewGetAPIInfoInternalServerError()
			}

			response := ConvertInfoResponse(res)
			return operations.NewGetAPIInfoOK().WithPayload(response)
		})
	}
	if api.PostAPIAuthHandler == nil {
		api.PostAPIAuthHandler = operations.PostAPIAuthHandlerFunc(func(params operations.PostAPIAuthParams) middleware.Responder {
			token, err := workers.Auth(*params.Body.Username, string(*params.Body.Password))
			if err != nil {
				return operations.NewPostAPIAuthBadRequest()
			}
			return operations.NewPostAPIAuthOK().WithPayload(&models.AuthResponse{Token: token})
		})
	}
	if api.PostAPISendCoinHandler == nil {
		api.PostAPISendCoinHandler = operations.PostAPISendCoinHandlerFunc(func(params operations.PostAPISendCoinParams, principal interface{}) middleware.Responder {
			token, err := auth.ExtractTokenFromHeader(params.HTTPRequest)
			if err != nil {
				return operations.NewGetAPIInfoUnauthorized()
			}
			err = workers.SendMoney(token, *params.Body.ToUser, uint64(*params.Body.Amount))
			if err != nil {
				return operations.NewGetAPIBuyItemInternalServerError()
			}
			return operations.NewGetAPIBuyItemOK()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	s.Addr = ":8080"
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	handler = auth.JWTAuthMiddleware(handler)

	// Другие middleware (например, CORS)
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func convertInventory(items *[]model.InventoryItem) []*models.InfoResponseInventoryItems0 {
	if items == nil {
		return nil
	}

	result := make([]*models.InfoResponseInventoryItems0, len(*items))
	for i, item := range *items {
		result[i] = &models.InfoResponseInventoryItems0{
			Quantity: item.Quantity,
			Type:     item.Type,
		}
	}
	return result
}

func convertCoins(coins uint64) int64 {
	return int64(coins)
}

func ConvertInfoResponse(res *model.InfoResponse) *models.InfoResponse {
	return &models.InfoResponse{
		Coins:       convertCoins(res.Coins),
		Inventory:   convertInventory(res.Inventory),
		CoinHistory: convertCoinHistory(res.CoinHistory),
	}
}

func convertCoinHistory(history model.CoinHistory) *models.InfoResponseCoinHistory {
	rec, send := make([]*models.InfoResponseCoinHistoryReceivedItems0, 0, len(*history.Received)), make([]*models.InfoResponseCoinHistorySentItems0, 0, len(*history.Sent))
	for _, val := range *history.Received {
		rec = append(rec, &models.InfoResponseCoinHistoryReceivedItems0{
			Amount:   val.Amount,
			FromUser: val.FromUser,
		})
	}
	for _, val := range *history.Sent {
		send = append(send, &models.InfoResponseCoinHistorySentItems0{
			Amount: val.Amount,
			ToUser: val.ToUser,
		})
	}
	return &models.InfoResponseCoinHistory{Sent: send, Received: rec}
}
