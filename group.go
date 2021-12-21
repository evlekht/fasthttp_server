package httpserver

import (
	"github.com/valyala/fasthttp"
)

type Group struct {
	middlewares []Middleware
	server      *HTTPServer
}

func (g *Group) Use(middlewares ...Middleware) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) POST(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	g.Handle("POST", path, handler, middlewares...)
}

func (g *Group) PUT(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	g.Handle("PUT", path, handler, middlewares...)
}

func (g *Group) GET(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	g.Handle("GET", path, handler, middlewares...)
}

func (g *Group) DELETE(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	g.Handle("DELETE", path, handler, middlewares...)
}

func (g *Group) PATCH(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	g.Handle("PATCH", path, handler, middlewares...)
}

func (g *Group) Handle(method, path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	m := make([]Middleware, 0, len(g.middlewares)+len(middlewares))
	m = append(m, g.middlewares...)
	m = append(m, middlewares...)
	g.server.Handle(method, path, handler, m...)
}
