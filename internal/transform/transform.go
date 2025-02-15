package transform

import "github.com/vsrtferrum/AvitoIntro/internal/model"

type BuyedItem struct {
	Item string `db:"name"`
}

type Transaction struct {
	Id         uint   `db:"id"`
	PersonName string `db:"name"`
	Value      uint64 `db:"cost"`
}
type NamePassword struct {
	Name     string `db:"name"`
	Password string `db:"password"`
}
type IdPassword struct {
	Id           uint64 `db:"id"`
	PasswordHash string `db:"password"`
}
type ItemIdCost struct {
	Id   uint64 `db:"id"`
	Cost uint64 `db:"cost"`
}

func (BuyedItem *BuyedItem) TransformBuyedItem() model.BuyedItem {
	return model.BuyedItem{Item: BuyedItem.Item}
}

func (Transaction *Transaction) TransformTransaction() model.Transaction {
	return model.Transaction{Id: Transaction.Id, PersonName: Transaction.PersonName, Value: Transaction.Value}
}

func (idPassword *IdPassword) TransformIdPassword() model.IdPassword {
	return model.IdPassword{Id: idPassword.Id, PasswordHash: idPassword.PasswordHash}
}

func (itemIdCost *ItemIdCost) TransformItemIdCost() model.ItemIdCost {
	return model.ItemIdCost{Id: itemIdCost.Id, Cost: itemIdCost.Cost}
}
