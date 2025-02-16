package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

type Cache struct {
	cache *redis.Client
	ttl   time.Duration
	CacheActions
}

func NewCahce(cacheModel model.CacheModel) (Cache, error) {
	cache := redis.NewClient(&redis.Options{
		Addr:     cacheModel.Addr,
		Password: "",
		DB:       0,
	})
	if cache == nil {
		return Cache{}, errors.ErrCacheCreation
	}
	ttl := time.Duration(cacheModel.TTL) * time.Second

	return Cache{cache: cache, ttl: ttl}, nil
}
