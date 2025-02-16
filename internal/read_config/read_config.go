package readconfig

import (
	"encoding/json"
	"os"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/transform"
)

func ReadConfig(path string) (model.CacheModel, string, string, int, model.WorkersModel, error) {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return model.CacheModel{}, "", "", 0, model.WorkersModel{}, errors.ErrParseConfig
	}
	return ParseConfig(jsonData)
}

func ParseConfig(jsonData []byte) (model.CacheModel, string, string, int, model.WorkersModel, error) {
	var config transform.Config
	err := json.Unmarshal(jsonData, &config)
	if err != nil {
		return model.CacheModel{}, "", "", 0, model.WorkersModel{}, errors.ErrParseConfig
	}
	cache, conn, logs, ttlJwt, workers := config.TransformConfig()
	return cache, conn, logs, ttlJwt, workers, nil
}
