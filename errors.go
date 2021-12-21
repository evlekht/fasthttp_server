package httpserver

import "errors"

var (
	errPortIsNotSet   = errors.New("port is not set for http server")
	errLoggerIsNotSet = errors.New("logger is not set for http server")
)
