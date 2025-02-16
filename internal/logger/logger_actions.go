package logger

import (
	"errors"

	"go.uber.org/zap"
)

type LoggerActions interface {
	WriteError(msg error) error
	WriteStatus(msg string) error
}

func (l *Logger) WriteError(msg error) {
	if l.errorLogger == nil {
		panic(errors.New("error logger is not initialized"))
	}

	l.errorLogger.Error("Error", zap.Error(msg))

}

func (l *Logger) WriteStatus(msg string) error {
	if l.infoLogger == nil {
		return errors.New("info logger is not initialized")
	}

	l.infoLogger.Info("Status", zap.String("message", msg))
	return nil
}
