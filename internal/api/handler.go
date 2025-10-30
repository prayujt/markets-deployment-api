package api

import (
	"log/slog"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	Log   *slog.Logger
	Redis *redis.Client
}

func NewHandler() *Handler {
	lvl := parseLevel(os.Getenv("LOGLEVEL"))
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	base := slog.New(h)

	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &Handler{Log: base, Redis: rd}
}

func parseLevel(s string) slog.Leveler {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
