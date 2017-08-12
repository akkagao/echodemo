package util

import "errors"

var (
	ERROR_SYSTEM      = errors.New("system error")
	ERROR_DATABASE    = errors.New("database error")
)
