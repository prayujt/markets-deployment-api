package services

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/event"

	"markets-api/internal/constants"
	"markets-api/internal/contracts"
	"markets-api/internal/database"
	"markets-api/internal/environment"
	"markets-api/internal/models"
)

func DeployMarket(env *environment.ServerEnvironment, market *models.MarketDeployRequest) error {
	tx, err := env.ContractHTTP.DeployMarket(env.ContractAuth, market.QuestionID, market.ConditionID, market.PositionIDYes, market.PositionIDNo)
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

func CompleteMarketDeployment(env *environment.ServerEnvironment, questionID string) {
	err := env.Queries.UpdateDeploymentStatus(context.Background(), database.UpdateDeploymentStatusParams{
		QuestionID: questionID,
		Status:     string(constants.MarketDeploymentStatusDeployed),
	})
	if err != nil {
		env.Log.Error("failed to update deployment status in database", "error", err)
		return
	}
	env.Log.Info("market deployment completed", "questionID", questionID)
}
