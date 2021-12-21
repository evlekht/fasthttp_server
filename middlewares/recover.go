package middlewares

import (
	"context"
	"fmt"

	httpserver "github.com/evlekht/fasthttp_server"
	"github.com/valyala/fasthttp"
)

func Recover(f func(ctx context.Context, err error)) httpserver.Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			defer func() {
				if err := recover(); err != nil {
					f(ctx, fmt.Errorf("%v", err))
				}
			}()

			next(ctx)
		}
	}
}
