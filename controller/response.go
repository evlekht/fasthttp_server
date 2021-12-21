package controller

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

func (c *controller) JSON(ctx *fasthttp.RequestCtx, statusCode int, response interface{}) {
	var body []byte
	var err error

	if marshallable, ok := response.(marshallable); ok {
		body, err = marshallable.MarshalJSON()
	} else {
		body, err = json.Marshal(response)
	}

	if err != nil {
		c.logger.Error(ctx, err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(body)
}

func (controller) Plain(ctx *fasthttp.RequestCtx, statusCode int, body []byte) {
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(body)
}

func (controller) NoContent(ctx *fasthttp.RequestCtx, statusCode int) {
	ctx.SetStatusCode(statusCode)
}
