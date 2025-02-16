package module

import (
	"context"
	"strconv"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/transform"
	"golang.org/x/crypto/bcrypt"
)

type ModuleActions interface {
	Auth(name, password string) (string, error)
	Buy(token string, itemName string) error
	Identify(token string) bool
	GetInfo(token string) (*[][]byte, error)
	SendMoney(token string, RecieverName string, amount uint64) error
}

func (module *Module) Auth(name, password string) (token string, err error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		module.logger.WriteError(err)
		return "", errors.ErrGenerateHash
	}

	data, err := module.database.Auth(model.AuthRequest{Username: name, Password: string(passwordHashed)})
	if err != nil {
		module.logger.WriteError(err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(passwordHashed, []byte(data.PasswordHash)); err != nil {
		module.logger.WriteError(err)
		return "", errors.ErrCompareHash
	}

	token, err = module.auth.GenerateFromPassword(data)
	return
}

func (module *Module) Identify(token string) bool {
	_, ok := module.auth.Identify(token)
	return ok
}

func (module *Module) Buy(token string, itemName string) (err error) {
	id, err := module.authByToken(token)
	if err != nil {
		return err
	}
	balance, err := module.database.GetUserBalanceById(id.Id)
	if err != nil {
		module.logger.WriteError(err)
		return errors.ErrExecQuery
	}

	item, err := module.database.GetItemCost(itemName)
	if err != nil {
		module.logger.WriteError(err)
		return errors.ErrExecQuery
	}

	if balance < item.Cost {
		module.logger.WriteError(errors.ErrSmallBalance)
		return errors.ErrSmallBalance
	}

	err = module.database.BuyItem(id.Id, item.Id, balance-item.Cost, item.Cost)
	if err != nil {
		module.logger.WriteError(err)
		return errors.ErrExecQuery
	}
	return
}

func (module *Module) GetInfo(token string) (*model.InfoResponse, error) {
	id, err := module.authByToken(token)
	if err != nil {
		return nil, err
	}
	cahceData, err := module.cache.Get(context.Background(), id.Id)
	if err == nil {
		data, err := transform.UnmarshalInfoResponse(cahceData)
		if err != nil {
			module.logger.WriteError(err)
		} else {
			return data, err
		}

	}
	data, err := module.database.TransactionsInfo(id.Id)
	if err != nil {
		module.logger.WriteError(err)
		return nil, err
	}

	marshalledData, err := transform.MarshalAllTransactionInfo(data, id.Username)
	if err != nil {
		module.logger.WriteError(err)
		return nil, err
	}

	err = module.cache.Set(context.Background(), strconv.FormatUint(id.Id, 10), marshalledData)
	if err != nil {
		module.logger.WriteError(err)
		module.logger.WriteError(errors.ErrSetValue)
	}
	return data, err
}
func (module *Module) SendMoney(token string, RecieverName string, amount uint64) error {
	idSender, err := module.authByToken(token)
	if err != nil {
		return err
	}
	balanceSender, err := module.database.GetUserBalanceById(idSender.Id)
	if err != nil {
		module.logger.WriteError(err)
		return errors.ErrExecQuery
	}

	if balanceSender < amount {
		module.logger.WriteError(errors.ErrSmallBalance)
		return errors.ErrSmallBalance
	}

	balanceReciever, err := module.database.GetUserBalanceByName(RecieverName)
	if err != nil {
		module.logger.WriteError(err)
		return errors.ErrExecQuery
	}

	err = module.database.SendMoney(idSender.Id, RecieverName, amount, balanceSender-amount, balanceReciever+amount)
	if err != nil {
		module.logger.WriteError(err)
		return errors.ErrExecQuery
	}
	return nil
}
