package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {
	logger := slog.New(&MiddlewareHandler{next: slog.NewJSONHandler(os.Stdout, nil)})
	slog.SetDefault(logger)

	ctx := context.Background()
	ctx = WithUserId(ctx, 123)

	now := time.Now()
	someHandler(ctx)
	slog.InfoContext(ctx, "someHandler", "duration", time.Since(now))
}

func someHandler(ctx context.Context) {
	// ctx = WithUserId(ctx, 456) // overwriting

	logger := slog.Default().With("child", "someHandler")
	logger.InfoContext(ctx, "msg")

	slog.InfoContext(ctx, "start")

	time.Sleep(time.Second)

	slog.InfoContext(ctx, "end")
}

type MiddlewareHandler struct {
	next slog.Handler
}

func (m *MiddlewareHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return m.next.Enabled(ctx, lvl)
}

type logCtxKey struct{}

type logCtx struct {
	userId int64
}

func WithUserId(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, logCtxKey{}, logCtx{userId: userId})
}

func (m *MiddlewareHandler) Handle(ctx context.Context, rec slog.Record) error {
	c, ok := ctx.Value(logCtxKey{}).(logCtx)
	if !ok {
		return m.next.Handle(ctx, rec)
	}

	if c.userId != 0 {
		rec.AddAttrs(slog.Int64("user_id", c.userId))
	}

	return m.next.Handle(ctx, rec)
}

func (m *MiddlewareHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MiddlewareHandler{next: m.next.WithAttrs(attrs)}
}

func (m *MiddlewareHandler) WithGroup(name string) slog.Handler {
	return &MiddlewareHandler{next: m.next.WithGroup(name)}
}
