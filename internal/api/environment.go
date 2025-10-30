package api

import (
	"log/slog"
	"os"
	"strings"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

type BaseEnvironment struct {
	Log *slog.Logger
}

type ClientEnvironment struct {
	BaseEnvironment
	QueueClient *asynq.Client
}

type ServerEnvironment struct {
	BaseEnvironment
	QueueServer *asynq.Server
}

func NewEnvironment() BaseEnvironment {
	lvl := parseLevel(os.Getenv("LOGLEVEL"))
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	base := slog.New(h)
	return BaseEnvironment{Log: base}
}

func NewClientEnvironment() *ClientEnvironment {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		// weird version issues, just disabling for now
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})

	client := asynq.NewClientFromRedisClient(redisClient)
	client.Ping()

	return &ClientEnvironment{BaseEnvironment: NewEnvironment(), QueueClient: client}
}

func NewServerEnvironment() *ServerEnvironment {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		// weird version issues, just disabling for now
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})

	srv := asynq.NewServerFromRedisClient(
		redisClient,
		asynq.Config{
			Concurrency: 10,
		},
	)
	srv.Ping()

	return &ServerEnvironment{BaseEnvironment: NewEnvironment(), QueueServer: srv}
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
