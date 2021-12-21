package controller

import "errors"

var (
	errLoggerIsNotSet = errors.New("logger is not set for http server")
)
