package main

import (
	"os"

	"github.com/hibiken/asynq"

	"markets-api/internal/api"
	"markets-api/internal/constants"
	"markets-api/internal/contracts"
	"markets-api/internal/services"
)

func main() {
	env := api.NewServerEnvironment()
	log := env.Log

	mux := asynq.NewServeMux()
	mux.HandleFunc(constants.MarketDeploymentQueueName, env.ProcessDeploymentRequest)

	go func() {
		log.Info("starting market deployment event listener")
		// cap buffer size at 128
		marketDeployments := make(chan *contracts.ContractsMarketDeployed, 128)
		sub, err := services.SubscribeMarketDeployments(log, env.ContractWS, marketDeployments)
		if err != nil {
			log.Error("could not subscribe to market deployments: %v", err)
			os.Exit(1)
		}
		defer sub.Unsubscribe()

		for event := range marketDeployments {
			log.Info("received MarketDeployed event", "questionID", event.QuestionId)
			services.CompleteMarketDeployment(log, env.Queries, event.QuestionId)
		}
	}()

	if err := env.QueueServer.Run(mux); err != nil {
		log.Error("could not run server: %v", err)
		os.Exit(1)
	}
}
