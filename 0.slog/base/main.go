package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

const (
	LevelNotice slog.Level = 1
)

func main() {
	// LEVELING
	fmt.Println("LEVELING")

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

	fmt.Println()

	// GROUPS
	fmt.Println("GROUPS")
	logger.InfoContext(ctx, "group", slog.Int64("int", 1), slog.Group("aboba", slog.String("a", "a"), slog.String("b", "b")))
	logger.Info("group")

	fmt.Println()

	// ENRICHED LOGGER
	fmt.Println("ENRICHED LOGGING")
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.With("app", "myapp", "version", "1.0.0").Info("Hello, world!")
	// logger.WithGroup("user").With("id", 123).Info("Hello, world!")

	// Можно вложить так чтобы все новые аттрибуты были частью группы
	logger1 := logger.With("func", 1).WithGroup("user").With("id", 123)
	f(context.WithValue(ctx, LoggerKey, logger1))

	// Можно вложить так чтобы группа была просто аттрибутом
	logger2 := logger.With("func", 2, "user", slog.GroupValue(slog.Int("id", 123)))
	f(context.WithValue(ctx, LoggerKey, logger2))
}

type LoggerCtxKey string

const (
	LoggerKey LoggerCtxKey = "logger"
)

func f(ctx context.Context) {
	logger := ctx.Value(LoggerKey).(*slog.Logger)
	logger.With("value", 1).Info("msg")
	logger.WithGroup("user").With("id", 1).Info("msg", "name", Name{"first", "last"})
}

// Для отображения структуры в логе
type Name struct {
	First, Last string
}

func (n Name) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("first", n.First),
		slog.String("last", n.Last),
	)
}

func LogExample(ctx context.Context, a int) error {
	logger := ctx.Value(LoggerKey).(*slog.Logger)

	// У каждого уровня логирования своя роль
	// INFO - отвечает за основные действия
	logger.Info("start")
	now := time.Now()

	var sum int
	for i := 0; i < a; i++ {
		// DEBUG - отвечает за детализированную информацию
		logger.Debug("step", "current_sum", sum, "increment", i)
		sum += i

		time.Sleep(100 * time.Millisecond)
	}

	if time.Since(now) > time.Second {
		// WARN - отвечает за некритические ошибки, на которые стоит обратить внимание
		logger.Warn("slow operation", "duration", time.Since(now))
	}

	logger.Info("finish", "sum", sum, "duration", time.Since(now))

	if sum > 100 {
		// ERROR - отвечает за ошибки
		err := fmt.Errorf("sum is too big")
		logger.Error("", "error", err)
		return err
	}

	return nil
}
