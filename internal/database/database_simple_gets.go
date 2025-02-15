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
