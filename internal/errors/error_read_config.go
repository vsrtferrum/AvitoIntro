package errors

import "errors"

var (
	ErrReadFromConfig = errors.New("read from config error")
	ErrParseConfig    = errors.New("parse config data error")
)
