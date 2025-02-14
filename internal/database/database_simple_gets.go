package database

import (
	"context"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/transform"
)

func (database *Database) listOfBuyedItems(id uint64) (*[]model.BuyedItem, error) {
	rows, err := database.pool.Query(context.Background(), listOfBuyedItems, id)

	if err != nil {
		database.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.BuyedItem, 0)
	var temp transform.BuyedItem
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

func (database *Database) getUserBalance(id uint64) (uint64, error) {
	rows, err := database.pool.Query(context.Background(), getUserBalance, id)
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

func (database *Database) getItemCost(itemName string) (uint64, error) {
	rows, err := database.pool.Query(context.Background(), getItemCost, itemName)
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

func (db *Database) sendedMoneyStat(id uint64) (*[]model.Transaction, error) {
	rows, err := db.pool.Query(context.Background(), sendedMoneyStat, id)

	if err != nil {
		db.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.Transaction, 0)
	var temp transform.Transaction
	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			db.log.WriteError(err)
			return nil, errors.ErrResultQuery
		}
		res = append(res, temp.TransformTransaction())
	}
	return &res, nil
}

func (db *Database) recievedMoneyStat(id uint64) (*[]model.Transaction, error) {
	rows, err := db.pool.Query(context.Background(), recievedMoneyStat, id)

	if err != nil {
		db.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.Transaction, 0)
	var temp transform.Transaction
	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			db.log.WriteError(err)
			return nil, errors.ErrResultQuery
		}
		res = append(res, temp.TransformTransaction())
	}
	return &res, nil
}
