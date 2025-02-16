package database

import (
	"context"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/transform"
)

type DatabaseActions interface {
	TransactionsInfo(id uint64) (*model.InfoResponse, error)
	Auth(name, password string) (model.IdPassword, error)
	BuyItem(idUser uint64, itemName string) error
	SendMoney(id uint64, toUser string, amount uint64) error
	GetUserBalanceById(id uint64) (uint64, error)
	GetUserBalanceByName(name string) (uint64, error)
	GetItemCost(itemName string) (model.ItemIdCost, error)
}

func (database *Database) Transaction(id uint64) (*model.InfoResponse, error) {
	balance, err := database.GetUserBalanceById(id)
	if err != nil {
		return nil, err
	}
	listOfBuyedItems, err := database.listOfBuyedItems(id)
	if err != nil {
		return nil, err
	}
	sendedMoney, err := database.sendedMoneyStat(id)
	if err != nil {
		return nil, err
	}
	recievedMoney, err := database.recievedMoneyStat(id)
	if err != nil {
		return nil, err
	}
	return &model.InfoResponse{Coins: balance, Inventory: listOfBuyedItems,
		CoinHistory: model.CoinHistory{Sent: sendedMoney, Received: recievedMoney}}, nil
}

func (database *Database) Auth(data model.AuthRequest) (model.IdPassword, error) {

	rows, err := database.pool.Query(context.Background(), authUser, data.Username, data.Password)
	if err != nil {
		database.log.WriteError(err)
		return model.IdPassword{}, errors.ErrResultQuery
	}
	res := make([]model.IdPassword, 0, 1)
	var temp transform.IdPassword
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return model.IdPassword{}, errors.ErrGetValue
		}
		res = append(res, temp.TransformIdPassword())
		if len(res) > 1 {
			database.log.WriteError(errors.ErrNonDeterministicUsers)
			return model.IdPassword{}, errors.ErrNonDeterministicUsers
		}
	}
	if len(res) == 0 {
		err := database.insertUser(data)
		if err != nil {
			database.log.WriteError(err)
			return model.IdPassword{}, errors.ErrResultQuery
		}
		return database.Auth(data)
	}
	return res[0], nil
}
func (database *Database) BuyItem(idUser, idItem uint64, newBalance, itemCost uint64) error {

	tx, err := database.pool.Begin(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCreateTransaction
	}
	_, err = tx.Exec(context.Background(), updateBalanceById, newBalance, idUser)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	_, err = tx.Exec(context.Background(), insertSale, idUser, idItem, itemCost)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	err = tx.Commit(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCommitTransaction
	}
	return nil
}

func (database *Database) SendMoney(idUser uint64, toUser string, amount, newBalanceSender, newBalanceReciever uint64) error {
	tx, err := database.pool.Begin(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCreateTransaction
	}
	_, err = tx.Exec(context.Background(), updateBalanceById, newBalanceSender, idUser)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	_, err = tx.Exec(context.Background(), updateBalanceByName, newBalanceReciever, toUser)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	_, err = tx.Exec(context.Background(), insertTransfer, idUser, toUser, amount)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	err = tx.Commit(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCommitTransaction
	}
	return nil
}

func (database *Database) GetUserBalanceById(id uint64) (uint64, error) {
	rows, err := database.pool.Query(context.Background(), getUserBalanceById, id)
	if err != nil {
		database.log.WriteError(err)
		return 0, errors.ErrResultQuery
	}

	res := make([]uint64, 0, 1)
	var temp uint64
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return 0, errors.ErrGetValue
		}

		res = append(res, temp)
		if len(res) > 1 {
			database.log.WriteError(errors.ErrNonDeterministicUsers)
			return 0, errors.ErrNonDeterministicUsers
		}
	}
	return res[0], nil
}
func (database *Database) GetUserBalanceByName(name string) (uint64, error) {
	rows, err := database.pool.Query(context.Background(), getUserBalanceByName, name)
	if err != nil {
		database.log.WriteError(err)
		return 0, errors.ErrResultQuery
	}

	res := make([]uint64, 0, 1)
	var temp uint64
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return 0, errors.ErrGetValue
		}
		res = append(res, temp)
		if len(res) > 1 {
			database.log.WriteError(errors.ErrNonDeterministicUsers)
			return 0, errors.ErrNonDeterministicUsers
		}
	}
	return res[0], nil
}

func (database *Database) GetItemCost(itemName string) (model.ItemIdCost, error) {
	rows, err := database.pool.Query(context.Background(), getItemCost, itemName)
	if err != nil {
		database.log.WriteError(err)
		return model.ItemIdCost{}, errors.ErrResultQuery
	}

	res := make([]model.ItemIdCost, 0, 1)
	var temp transform.ItemIdCost
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return model.ItemIdCost{}, errors.ErrGetValue
		}

		res = append(res, temp.TransformItemIdCost())
		if len(res) > 1 {
			database.log.WriteError(errors.ErrNonDeterministicUsers)
			return model.ItemIdCost{}, errors.ErrNonDeterministicUsers
		}
	}
	return res[0], nil
}
