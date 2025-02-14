package errors

import "errors"

var (
	ErrConnectionTimeLimit = errors.New("connection time limit")
	ErrCreateConnection    = errors.New("failed to create connection")
	ErrCreateConfig        = errors.New("failed to create config")
	ErrCloseConnection     = errors.New("error while closing connection")
	ErrSendQuery           = errors.New("error sending query or error responce")
	ErrConvertResponce     = errors.New("error converting database responce")
	ErrResultQuery         = errors.New("error nil result")
)
