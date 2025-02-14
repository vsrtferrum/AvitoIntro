package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
)

type CacheActions interface {
	GetListOfBuyedItems(cntx context.Context, id uint64) (*[][]byte, error)
	GetSendedMoneyStat(cntx context.Context, id uint64) (*[][]byte, error)
	GetRecievedMoneyStat(cntx context.Context, id uint64) (*[][]byte, error)
	SetListOfBuyedItems(cntx context.Context, query string, data *[][]byte) error
	SetSendedMoneyStat(cntx context.Context, query string, data *[][]byte) error
	SetRecievedMoneyStat(cntx context.Context, query string, data *[][]byte) error
}

func (cache *Cache) GetListOfBuyedItems(cntx context.Context, id uint64) (*[][]byte, error) {
	data, err := cache.CacheBuyedIems.Get(context.Background(), strconv.FormatUint(id, 10)).Result()
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

func (cache *Cache) GetSendedMoneyStat(cntx context.Context, id uint64) (*[][]byte, error) {
	data, err := cache.CacheSendedMoneyStat.Get(context.Background(), strconv.FormatUint(id, 10)).Result()
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

func (cache *Cache) GetRecievedMoneyStat(cntx context.Context, id uint64) (*[][]byte, error) {
	data, err := cache.CacheRecievedMoneyStat.Get(context.Background(), strconv.FormatUint(id, 10)).Result()
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

func (cache *Cache) SetListOfBuyedItems(cntx context.Context, query string, data *[][]byte) error {
	temp, err := json.Marshal(data)
	if err != nil {
		return errors.ErrJsonMarshall
	}
	return cache.CacheBuyedIems.Set(cntx, query, temp, 0).Err()
}

func (cache *Cache) SetSendedMoneyStat(cntx context.Context, query string, data *[][]byte) error {
	temp, err := json.Marshal(data)
	if err != nil {
		return errors.ErrJsonMarshall
	}
	return cache.CacheSendedMoneyStat.Set(cntx, query, temp, 0).Err()
}

func (cache *Cache) SetRecievedMoneyStat(cntx context.Context, query string, data *[][]byte) error {
	temp, err := json.Marshal(data)
	if err != nil {
		return errors.ErrJsonMarshall
	}
	return cache.CacheRecievedMoneyStat.Set(cntx, query, temp, 0).Err()
}
