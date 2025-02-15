package module

import (
	"github.com/vsrtferrum/AvitoIntro/internal/auth"
	"github.com/vsrtferrum/AvitoIntro/internal/cache"
	"github.com/vsrtferrum/AvitoIntro/internal/database"
	"github.com/vsrtferrum/AvitoIntro/internal/logger"
)

type Module struct {
	cache    *cache.Cache
	database *database.Database
	logger   *logger.Logger
	auth     *auth.Auth
	ModuleActions
}

func NewModule(cache *cache.Cache, database *database.Database, logger *logger.Logger, auth *auth.Auth) Module {
	return Module{cache: cache, database: database, logger: logger, auth: auth}
}
