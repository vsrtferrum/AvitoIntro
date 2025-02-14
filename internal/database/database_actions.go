package database

import (
	"context"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/transform"
)

type DatabaseActions interface {
	ListOfBuyedItems(id uint64) (*[]model.BuyedItem, error)
	SendedMoneyStat(id uint64) (*[]model.Transaction, error)
	RecievedMoneyStat(id uint64) (*[]model.Transaction, error)
}

func (db *Database) ListOfBuyedItems(id uint64) (*[]model.BuyedItem, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT shop.name"+
		"FROM sales"+
		"JOIN shop ON sales.item_id = shop.id"+
		"WHERE id = $1", id)

	if err != nil {
		db.log.WriteError(err)
		return nil, errors.ErrResultQuery
	}
	defer rows.Close()

	res := make([]model.BuyedItem, 0)
	var temp transform.BuyedItem
	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			db.log.WriteError(err)
			return nil, errors.ErrResultQuery
		}
		res = append(res, temp.TransformBuyedItem())
	}
	return &res, nil
}

func (db *Database) SendedMoneyStat(id uint64) (*[]model.Transaction, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT users.name, users.id, transations.cost"+
		"FROM users"+
		"JOIN transations ON transations.sender_id = users.id"+
		"WHERE id = $1", id)

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

func (db *Database) RecievedMoneyStat(id uint64) (*[]model.Transaction, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT users.name, users.id, transations.cost"+
		"FROM users"+
		"JOIN transations ON transations.recipient_id = users.id"+
		"WHERE id = $1", id)

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
