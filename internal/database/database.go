package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vsrtferrum/AvitoIntro/internal/logger"
)

type Database struct {
	connStr string
	pool    *pgxpool.Pool
	log     *logger.Logger
	DatabaseConnection
	DatabaseActions
}

func NewDatabase(connStr string) Database {
	return Database{connStr: connStr}
}
