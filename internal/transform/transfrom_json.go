package transform

import "github.com/vsrtferrum/AvitoIntro/internal/model"

type InfoResponse struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinHistory struct {
	Received []CoinTransaction `json:"received"`
	Sent     []CoinTransaction `json:"sent"`
}

type CoinTransaction struct {
	FromUser string `json:"fromUser,omitempty"`
	ToUser   string `json:"toUser,omitempty"`
	Amount   int    `json:"amount"`
}

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type SendCoinRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type Config struct {
	ConnStr         string `json:"connStr"`
	LogPath         string `json:"logPath"`
	CacheAddr       string `json:"cacheAddr"`
	TTL             int    `json:"TTL"`
	TTLJWT          int    `json:"TTLJWT"`
	WorkersCount    int    `json:"workersCount"`
	WorkersQueueLen int    `json:"workersQueueLen"`
}

func (config *Config) TransformConfig() (model.CacheModel, string, string, int, model.WorkersModel) {
	return model.CacheModel{Addr: config.CacheAddr, TTL: config.TTL}, config.ConnStr, config.LogPath, config.TTLJWT, model.WorkersModel{WorkersCount: config.TTLJWT, WorkersQueueLen: config.WorkersQueueLen}
}
