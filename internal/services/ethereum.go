package services

import (
	"context"
	"log/slog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/event"

	"markets-api/internal/constants"
	"markets-api/internal/contracts"
	"markets-api/internal/database"
	"markets-api/internal/models"
)

func DeployMarket(log *slog.Logger, contract *contracts.Contracts, auth *bind.TransactOpts, market *models.MarketDeployRequest) error {
	tx, err := contract.DeployMarket(auth, market.QuestionID, market.ConditionID, market.PositionIDYes, market.PositionIDNo)
	if err != nil {
		log.Error("failed to deploy market contract", "error", err)
		return err
	}
	log.Info("market deployment transaction sent", "txHash", tx.Hash().Hex())
	return nil
}

func SubscribeMarketDeployments(log *slog.Logger, contractWS *contracts.Contracts, events chan *contracts.ContractsMarketDeployed) (event.Subscription, error) {
	sub, err := contractWS.WatchMarketDeployed(&bind.WatchOpts{Context: context.Background()}, events, nil)
	if err != nil {
		log.Error("failed to subscribe to MarketDeployed events", "error", err)
		return nil, err
	}
	return sub, nil
}

func CompleteMarketDeployment(log *slog.Logger, queries *database.Queries, questionID string) {
	err := queries.UpdateDeploymentStatus(context.Background(), database.UpdateDeploymentStatusParams{
		QuestionID: questionID,
		Status:     string(constants.MarketDeploymentStatusDeployed),
	})
	if err != nil {
		log.Error("failed to update deployment status in database", "error", err)
		return
	}
	log.Info("market deployment completed", "questionID", questionID)
}
