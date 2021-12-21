package controller

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func (c *controller) UnmarshalBody(ctx *fasthttp.RequestCtx, target interface{}) error {
	if unmarshallable, ok := target.(unmarshallable); ok {
		if err := unmarshallable.UnmarshalJSON(ctx.PostBody()); err != nil {
			return err
		}
	} else {
		if err := json.Unmarshal(ctx.PostBody(), target); err != nil {
			return err
		}
	}

	return nil
}

func (c *controller) PathParam(ctx *fasthttp.RequestCtx, key string) string {
	userValue := ctx.UserValue(key)
	if userValue == nil {
		return ""
	}

	param, ok := userValue.(string)
	if !ok {
		return ""
	}

	return param
}

func (c *controller) QueryParam(ctx *fasthttp.RequestCtx, key string) string {
	return string(ctx.QueryArgs().Peek(key))
}
