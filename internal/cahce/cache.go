package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	readconfig "github.com/vsrtferrum/AvitoIntro/internal/read_config"
)

type Cache struct {
	CacheBuyedIems, CacheSendedMoneyStat, CacheRecievedMoneyStat *redis.Client
	CacheActions
}

func NewCahce(cacheModel readconfig.CacheModel) (Cache, error) {
	CacheBuyedIems := redis.NewClient(&redis.Options{
		Addr:     cacheModel.AddrBuyedIems,
		Password: "",
		DB:       0,
	})
	if CacheBuyedIems == nil {
		return Cache{}, errors.ErrCacheCreation
	}

	CacheSendedMoneyStat := redis.NewClient(&redis.Options{
		Addr:     cacheModel.AddrSendedMoneyStat,
		Password: "",
		DB:       0,
	})
	if CacheSendedMoneyStat == nil {
		return Cache{}, errors.ErrCacheCreation
	}

	CacheRecievedMoneyStat := redis.NewClient(&redis.Options{
		Addr:     cacheModel.AddrRecievedMoneyStat,
		Password: "",
		DB:       0,
	})
	if CacheRecievedMoneyStat == nil {
		return Cache{}, errors.ErrCacheCreation
	}
	return Cache{CacheBuyedIems: CacheBuyedIems,
		CacheSendedMoneyStat:   CacheSendedMoneyStat,
		CacheRecievedMoneyStat: CacheRecievedMoneyStat}, nil
}
