package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/vsrtferrum/AvitoIntro/internal/errors"
)

type CacheActions interface {
	Get(cntx context.Context, id uint64) (*[]byte, error)
	Set(cntx context.Context, query string, data *[]byte) error
}

func (cache *Cache) Get(cntx context.Context, id uint64) (*[]byte, error) {
	data, err := cache.cache.Get(cntx, strconv.FormatUint(id, 10)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.ErrGetValue
		}
		return nil, errors.ErrGetValue
	}

	result := []byte(data)
	return &result, nil
}

func (cache *Cache) Set(cntx context.Context, id string, data *[]byte) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.ErrJsonMarshall
	}

	return cache.cache.Set(cntx, id, jsonData, cache.ttl).Err()
}
