package errors

import "errors"

var (
	ErrCacheCreation = errors.New("failed to create cache")
	ErrGetValue      = errors.New("failed get data from cahce")
	ErrSetValue      = errors.New("failed to set data")
)
