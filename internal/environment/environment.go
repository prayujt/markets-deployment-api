package environment

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"

	"markets-api/internal/contracts"
	"markets-api/internal/database"
	"markets-api/internal/utils"
)

type BaseEnvironment struct {
	Log         *slog.Logger
	RedisClient *redis.Client
	Queries     *database.Queries
}

type ClientEnvironment struct {
	BaseEnvironment
	QueueClient *asynq.Client
}

type ServerEnvironment struct {
	BaseEnvironment
	QuestionIDSet *utils.Set
	QueueServer   *asynq.Server
	ContractHTTP  *contracts.Contracts
	ContractWS    *contracts.Contracts
	ContractAuth  *bind.TransactOpts
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
		base.Error("failed to open DB connection", "error", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		base.Error("failed to ping DB", "error", err)
		os.Exit(1)
	}
	queries := database.New(db)

	return BaseEnvironment{Log: base, RedisClient: redisClient, Queries: queries}
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

	httpURL := os.Getenv("ETH_HTTP_RPC")
	wsURL := os.Getenv("ETH_WS_RPC")
	contractAddress := os.Getenv("MARKETS_CONTRACT")
	privateKey := os.Getenv("ETH_PRIVATE_KEY")
	if httpURL == "" || wsURL == "" || contractAddress == "" || privateKey == "" {
		env.Log.Error("missing required ETH environment variables")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	httpClient, err := ethclient.DialContext(ctx, httpURL)
	if err != nil {
		env.Log.Error("eth http dial failed", "err", err)
		os.Exit(1)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	wsClient, err := ethclient.DialContext(ctx, wsURL)
	if err != nil {
		env.Log.Error("eth ws dial failed", "err", err)
		os.Exit(1)
	}

	contractAddr := common.HexToAddress(contractAddress)
	contractHTTP, err := contracts.NewContracts(contractAddr, httpClient)
	if err != nil {
		env.Log.Error("failed http contract", "error", err)
		os.Exit(1)
	}
	contractWS, err := contracts.NewContracts(contractAddr, wsClient)
	if err != nil {
		env.Log.Error("failed ws contract", "error", err)
		os.Exit(1)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	chainID, cidErr := httpClient.ChainID(ctx)
	cancel()
	var contractAuth *bind.TransactOpts
	if cidErr != nil {
		env.Log.Error("get chain id failed", "err", cidErr)
	} else {
		key, keyErr := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
		if keyErr != nil {
			env.Log.Error("bad private key", "err", keyErr)
		} else {
			contractAuth, err = bind.NewKeyedTransactorWithChainID(key, chainID)
			if err != nil {
				env.Log.Error("transactor create failed", "err", err)
			}
		}
	}

	return &ServerEnvironment{BaseEnvironment: env, QueueServer: srv, ContractHTTP: contractHTTP, ContractWS: contractWS, ContractAuth: contractAuth, QuestionIDSet: utils.NewSet()}
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
