package transform

import (
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

type InventoryItemDB struct {
	Type     string `db:"name"`
	Quantity int64  `db:"item_count"`
}
type CoinTransactionDB struct {
	FromUser string `db:"sender_name"`
	ToUser   string `db:"recipient_name"`
	Amount   int64  `db:"cost"`
}
type AuthRequestDB struct {
	Username string `db:"name"`
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

func (InventoryItemDB *InventoryItemDB) TransformBuyedItem() model.InventoryItem {
	return model.InventoryItem{Type: InventoryItemDB.Type, Quantity: InventoryItemDB.Quantity}
}

func (CoinTransactionDB *CoinTransactionDB) TransformTransaction() model.CoinTransaction {
	return model.CoinTransaction{FromUser: CoinTransactionDB.FromUser, ToUser: CoinTransactionDB.ToUser, Amount: CoinTransactionDB.Amount}
}

func (AuthRequestDB *AuthRequestDB) TransformIdPassword() model.AuthRequest {
	return model.AuthRequest{Username: AuthRequestDB.Username, Password: AuthRequestDB.Password}
}

func (idPassword *IdPassword) TransformIdPassword() model.IdPassword {
	return model.IdPassword{Id: idPassword.Id, PasswordHash: idPassword.PasswordHash}
}

func (itemIdCost *ItemIdCost) TransformItemIdCost() model.ItemIdCost {
	return model.ItemIdCost{Id: itemIdCost.Id, Cost: itemIdCost.Cost}
}
