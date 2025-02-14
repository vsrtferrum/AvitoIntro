package module

import (
	"github.com/vsrtferrum/AvitoIntro/internal/cache"
	"github.com/vsrtferrum/AvitoIntro/internal/database"
	"github.com/vsrtferrum/AvitoIntro/internal/logger"
)

type Module struct {
	cache    *cache.Cache
	database *database.Database
	logger   *logger.Logger
	ModuleActions
}

func NewModule(cache *cache.Cache, database *database.Database, logger *logger.Logger) Module {
	return Module{cache: cache, database: database, logger: logger}
}
