package errors

import "errors"

var (
	ErrAuthTokenRequired = errors.New("error no auth token")
	ErrAuthTokenFormat   = errors.New("error wrong token format")
)
