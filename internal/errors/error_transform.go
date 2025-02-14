package errors

import "errors"

var (
	ErrJsonMarshall   = errors.New("marshall json error")
	ErrJsonUnMarshall = errors.New("unmarshall json error")
)
