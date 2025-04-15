package main

import (
	"context"
	"log/slog"
	"os"
)

const (
	LevelNotice slog.Level = 1
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey && a.Value.Equal(slog.AnyValue(LevelNotice)) { // add notice level
				return slog.Attr{
					Key:   a.Key,
					Value: slog.StringValue("NOTICE"),
				}
			}
			return a
		},
	}))
	logger.Debug("Hello, debug!")
	logger.Info("Hello, info!")
	logger.Warn("Hello, warn!")
	logger.Error("Hello, error!")

	ctx := context.Background()
	logger.LogAttrs(ctx, LevelNotice, "Hello, notice!")
}
