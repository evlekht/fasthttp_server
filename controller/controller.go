package controller

import (
	httpserver "github.com/evlekht/fasthttp_server"
	"github.com/valyala/fasthttp"
)

type unmarshallable interface {
	UnmarshalJSON([]byte) error
}
type marshallable interface {
	MarshalJSON() ([]byte, error)
}

type Controller interface {
	JSON(ctx *fasthttp.RequestCtx, statusCode int, response interface{})
	Plain(ctx *fasthttp.RequestCtx, statusCode int, body []byte)
	NoContent(ctx *fasthttp.RequestCtx, statusCode int)

	UnmarshalBody(ctx *fasthttp.RequestCtx, target interface{}) error
	PathParam(ctx *fasthttp.RequestCtx, key string) string
	QueryParam(ctx *fasthttp.RequestCtx, key string) string
}

type controller struct {
	logger httpserver.Logger
}

func NewController(logger httpserver.Logger) (Controller, error) {
	if logger == nil {
		return nil, errLoggerIsNotSet
	}

	controller := &controller{
		logger: logger,
	}

	return controller, nil
}
