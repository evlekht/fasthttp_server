package httpserver

import (
	"context"
)

type Logger interface {
	Error(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
}
