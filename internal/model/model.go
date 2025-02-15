package model

type BuyedItem struct {
	Item string
}
type Transaction struct {
	Id         uint
	PersonName string
	Value      uint64
}
type IdPassword struct {
	Id           uint64
	PasswordHash string
}
type ItemIdCost struct {
	Id, Cost uint64
}
type AllTransactionInfo struct {
	BuyedItem     *[]BuyedItem
	SendedMoney   *[]Transaction
	RecievedMoney *[]Transaction
}
type NamePassword struct {
	Name, Password string
}
