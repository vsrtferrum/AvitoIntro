package logger

import (
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	infoLogger  *zap.Logger
	errorLogger *zap.Logger
	LoggerActions
}

func NewLogger(logPath string) (*Logger, error) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	infoCfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout", logPath + "/info_" + timestamp + ".json"},
	}

	errorCfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.ErrorLevel),
		OutputPaths: []string{"stdout", logPath + "/errors_" + timestamp + ".json"},
	}

	infoLogger, err := infoCfg.Build()
	if err != nil {
		return nil, err
	}

	errorLogger, err := errorCfg.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}, nil
}
