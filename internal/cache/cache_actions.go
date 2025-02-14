package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
)

type CacheActions interface {
	Get(cntx context.Context, id uint64) (*[][]byte, error)
	Set(cntx context.Context, query string, data *[][]byte) error
}

func (cache *Cache) Get(cntx context.Context, id uint64) (*[][]byte, error) {
	data, err := cache.cache.Get(context.Background(), strconv.FormatUint(id, 10)).Result()
	if err != nil {
		return nil, errors.ErrGetValue
	}
	var resJson [][]byte
	err = json.Unmarshal([]byte(data), &resJson)
	if err != nil {
		return nil, errors.ErrJsonMarshall
	}
	return &resJson, nil
}

func (cache *Cache) Set(cntx context.Context, query string, data *[][]byte) error {
	temp, err := json.Marshal(data)
	if err != nil {
		return errors.ErrJsonMarshall
	}
	return cache.cache.Set(cntx, query, temp, cache.ttl).Err()
}
