package logger

import (
	"errors"

	"go.uber.org/zap"
)

type LoggerActions interface {
	WriteError(msg error) error
	WriteStatus(msg string) error
}

func (l *Logger) WriteError(msg error) error {
	if l.errorLogger == nil {
		return errors.New("error logger is not initialized")
	}

	l.errorLogger.Error("Error", zap.Error(msg))
	return nil
}

func (l *Logger) WriteStatus(msg string) error {
	if l.infoLogger == nil {
		return errors.New("info logger is not initialized")
	}

	l.infoLogger.Info("Status", zap.String("message", msg))
	return nil
}
