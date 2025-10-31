package services

import (
	"context"
	"database/sql"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/event"
	"github.com/sqlc-dev/pqtype"

	"markets-api/internal/contracts"
	"markets-api/internal/database"
	"markets-api/internal/environment"
	"markets-api/internal/models"
)

func DeployMarket(env *environment.ServerEnvironment, market *models.MarketDeployRequest) error {
	tx, err := env.ContractHTTP.DeployMarket(env.ContractAuth, "", "", "", "")
	if err != nil {
		env.Log.Error("failed to deploy market contract", "error", err)
		return err
	}
	env.Log.Info("market deployment transaction sent", "txHash", tx.Hash().Hex())
	return nil
}

func SubscribeMarketDeployments(env *environment.ServerEnvironment, events chan *contracts.ContractsMarketDeployed) (event.Subscription, error) {
	sub, err := env.ContractWS.WatchMarketDeployed(&bind.WatchOpts{Context: context.Background()}, events, nil)
	if err != nil {
		env.Log.Error("failed to subscribe to MarketDeployed events", "error", err)
		return nil, err
	}
	return sub, nil
}

func CompleteMarketDeployment(env *environment.ServerEnvironment, marketID string) {
	err := env.Queries.SetMarketDeployed(context.Background(), database.SetMarketDeployedParams{
		MarketID:    marketID,
		QuestionID:  sql.NullString{Valid: true, String: "sample_question_id"},
		ConditionID: sql.NullString{Valid: true, String: "sample_condition_id"},
		ClobTokenIds: pqtype.NullRawMessage{
			Valid:      true,
			RawMessage: []byte(`["token1","token2"]`),
		},
	})
	if err != nil {
		env.Log.Error("failed to update deployment status in database", "error", err)
		return
	}
	env.Log.Info("market deployment completed", "questionID", marketID)
}
