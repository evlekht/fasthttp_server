package httpserver

import (
	"context"
	"fmt"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Middleware func(next fasthttp.RequestHandler) fasthttp.RequestHandler

type HTTPServer struct {
	middlewares []Middleware
	router      *router.Router
	httpserver  *fasthttp.Server
	logger      Logger
	port        string
}

func New(logger Logger, port string, middlewares ...Middleware) (*HTTPServer, error) {
	if port == "" {
		return nil, errPortIsNotSet
	}
	if logger == nil {
		return nil, errLoggerIsNotSet
	}

	server := &HTTPServer{
		logger:      logger,
		router:      router.New(),
		port:        port,
		middlewares: middlewares,
	}

	server.httpserver = &fasthttp.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server, nil
}

func (s *HTTPServer) Start(ctx context.Context) error {
	s.httpserver.Handler = applyMiddlewares(s.router.Handler, s.middlewares...)
	s.logger.Info(ctx, fmt.Sprintf("HTTP is listening on %s", s.port))
	return s.httpserver.ListenAndServe(":" + s.port)
}

func (s HTTPServer) Shutdown(ctx context.Context) error {
	if err := s.httpserver.Shutdown(); err != nil {
		s.logger.Error(ctx, err)
		return nil
	}

	s.logger.Info(ctx, fmt.Sprintf("HTTP have stopped listening on %s", s.port))

	return nil
}

func (s *HTTPServer) Use(middlewares ...Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}

func (s *HTTPServer) POST(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	s.Handle("POST", path, handler, middlewares...)
}

func (s *HTTPServer) PUT(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	s.Handle("PUT", path, handler, middlewares...)
}

func (s *HTTPServer) GET(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	s.Handle("GET", path, handler, middlewares...)
}

func (s *HTTPServer) DELETE(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	s.Handle("DELETE", path, handler, middlewares...)
}

func (s *HTTPServer) PATCH(path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	s.Handle("PATCH", path, handler, middlewares...)
}

func (s *HTTPServer) Handle(method, path string, handler fasthttp.RequestHandler, middlewares ...Middleware) {
	s.router.Handle(method, path, applyMiddlewares(handler, middlewares...))
}

func applyMiddlewares(handler fasthttp.RequestHandler, middlewares ...Middleware) fasthttp.RequestHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
