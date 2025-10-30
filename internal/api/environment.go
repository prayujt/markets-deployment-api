package api

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"

	"markets-api/internal/database"
)

type BaseEnvironment struct {
	Log         *slog.Logger
	RedisClient *redis.Client
}

type ClientEnvironment struct {
	BaseEnvironment
	QueueClient *asynq.Client
}

type ServerEnvironment struct {
	BaseEnvironment
	QueueServer *asynq.Server
	Queries     *database.Queries
}

func NewEnvironment() BaseEnvironment {
	lvl := parseLevel(os.Getenv("LOGLEVEL"))
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	base := slog.New(h)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		// weird version issues, just disabling for now
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})

	return BaseEnvironment{Log: base, RedisClient: redisClient}
}

func NewClientEnvironment() *ClientEnvironment {
	env := NewEnvironment()
	client := asynq.NewClientFromRedisClient(env.RedisClient)
	client.Ping()

	return &ClientEnvironment{BaseEnvironment: env, QueueClient: client}
}

func NewServerEnvironment() *ServerEnvironment {
	env := NewEnvironment()
	srv := asynq.NewServerFromRedisClient(
		env.RedisClient,
		asynq.Config{
			Concurrency: 10,
		},
	)
	srv.Ping()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("PSQL_USERNAME"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_DATABASE"),
	)

	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		env.Log.Error("failed to open DB connection", "error", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		env.Log.Error("failed to ping DB", "error", err)
		os.Exit(1)
	}

	queries := database.New(db)

	return &ServerEnvironment{BaseEnvironment: env, QueueServer: srv, Queries: queries}
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
