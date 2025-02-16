package model

type InfoResponse struct {
	Coins       uint64
	Inventory   *[]InventoryItem
	CoinHistory CoinHistory
}

type InventoryItem struct {
	Type     string
	Quantity int64
}

type CoinHistory struct {
	Received *[]CoinTransaction
	Sent     *[]CoinTransaction
}
type AuthAns struct {
	Id       uint64
	Username string
}

type IdPassword struct {
	Id           uint64
	PasswordHash string
	Username     string
}
type ItemIdCost struct {
	Id, Cost uint64
}

type CoinTransaction struct {
	FromUser string
	ToUser   string
	Amount   int64
}

type ErrorResponse struct {
	Errors string
}

type AuthRequest struct {
	Username string
	Password string
}

type AuthResponse struct {
	Token string
}

type SendCoinRequest struct {
	ToUser string
	Amount int
}

type CacheModel struct {
	Addr string
	TTL  int
}
type WorkersModel struct {
	WorkersCount, WorkersQueueLen int
}
