package logger

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	//logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	//	Level: slog.LevelDebug,
	//}))

	slog.SetDefault(logger)
	return logger
}
