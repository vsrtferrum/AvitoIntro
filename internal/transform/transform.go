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

func (BuyedItem *BuyedItem) TransformBuyedItem() model.BuyedItem {
	return model.BuyedItem{Item: BuyedItem.Item}
}

func (Transaction *Transaction) TransformTransaction() model.Transaction {
	return model.Transaction{Id: Transaction.Id, PersonName: Transaction.PersonName, Value: Transaction.Value}
}
