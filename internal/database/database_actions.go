package database

import (
	"context"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

type DatabaseActions interface {
	TransactionsInfo(id uint64) (*[]model.AllTransactionInfo, error)
	Auth(name, password string) (bool, error)
	BuyItem(idUser uint64, itemName string) error
	SendMoney(id uint64, toUser string, amount uint64) error
}

func (database *Database) Transaction(id uint64) (*model.AllTransactionInfo, error) {
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
	return &model.AllTransactionInfo{BuyedItem: listOfBuyedItems, SendedMoney: sendedMoney,
		RecievedMoney: recievedMoney}, nil
}

func (database *Database) Auth(data model.NamePassword) (uint64, error) {

	rows, err := database.pool.Query(context.Background(), authUser, data.Name, data.Password)
	if err != nil {
		database.log.WriteError(err)
		return 0, errors.ErrResultQuery
	}
	res := make([]uint64, 0, 1)
	var temp uint64
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&temp)
		res = append(res, temp)
		if len(res) > 1 {
			database.log.WriteError(errors.ErrNonDeterministicUsers)
			return 0, errors.ErrNonDeterministicUsers
		}
	}
	return res[0], nil
}
func (database *Database) BuyItem(idUser uint64, itemName string) error {
	balance, err := database.getUserBalance(idUser)
	if err != nil {
		return err
	}

	itemCost, err := database.getItemCost(itemName)
	if err != nil {
		return err
	}

	if balance < itemCost {
		return errors.ErrSmallBalance
	}

	tx, err := database.pool.Begin(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCreateTransaction
	}
	_, err = tx.Exec(context.Background(), updateBalance, balance-itemCost, idUser)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	_, err = tx.Exec(context.Background(), insertSale, idUser, itemName, itemCost)
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

func (database *Database) SendMoney(idUser uint64, toUser string, amount uint64) error {
	balanceSender, err := database.getUserBalance(idUser)
	if err != nil {
		return err
	}

	if balanceSender < amount {
		database.log.WriteError(errors.ErrSmallBalance)
		return errors.ErrSmallBalance
	}
	balanceReciever, err := database.getUserBalance(idUser)
	if err != nil {
		return err
	}
	tx, err := database.pool.Begin(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCreateTransaction
	}
	_, err = tx.Exec(context.Background(), updateBalance, balanceSender-amount, idUser)
	if err != nil {
		database.log.WriteError(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			database.log.WriteError(err)
		}
		return errors.ErrExecTransaction
	}

	_, err = tx.Exec(context.Background(), updateBalance, balanceReciever+amount, toUser)
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
