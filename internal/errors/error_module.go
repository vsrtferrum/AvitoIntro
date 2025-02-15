package errors

import "errors"

var (
	ErrGenerateHash = errors.New("error generate hash")
	ErrCompareHash  = errors.New("error compare hash")
	ErrNoUserFound  = errors.New("error user not found")
	ErrExecQuery    = errors.New("error exec query")
	ErrSmallBalance = errors.New("error balance size")
)
