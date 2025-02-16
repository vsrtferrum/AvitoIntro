package database

import (
	"context"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/transform"
)

func (database *Database) listOfBuyedItems(id uint64) (*[]model.InventoryItem, error) {
	rows, err := database.pool.Query(context.Background(), listOfBuyedItems, id)

	if err != nil {
		database.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.InventoryItem, 0)
	var temp transform.InventoryItemDB
	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return nil, errors.ErrResultQuery
		}
		res = append(res, temp.TransformBuyedItem())
	}
	return &res, nil
}

func (database *Database) sendedMoneyStat(id uint64) (*[]model.CoinTransaction, error) {
	rows, err := database.pool.Query(context.Background(), sendedMoneyStat, id)

	if err != nil {
		database.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.CoinTransaction, 0)
	var temp transform.CoinTransactionDB
	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return nil, errors.ErrResultQuery
		}
		res = append(res, temp.TransformTransaction())
	}
	return &res, nil
}

func (database *Database) recievedMoneyStat(id uint64) (*[]model.CoinTransaction, error) {
	rows, err := database.pool.Query(context.Background(), recievedMoneyStat, id)

	if err != nil {
		database.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.CoinTransaction, 0)
	var temp transform.CoinTransactionDB
	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			database.log.WriteError(err)
			return nil, errors.ErrResultQuery
		}
		res = append(res, temp.TransformTransaction())
	}
	return &res, nil
}

func (database *Database) insertUser(user model.AuthRequest) error {
	tx, err := database.pool.Begin(context.Background())
	if err != nil {
		database.log.WriteError(err)
		return errors.ErrCreateTransaction
	}

	_, err = tx.Exec(context.Background(), insertUser, user.Username, user.Password)
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
