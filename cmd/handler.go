package main

import (
	"os"

	"github.com/hibiken/asynq"

	"markets-api/internal/constants"
	"markets-api/internal/handlers"
)

func main() {
	env := handlers.NewServerEnvironment()
	log := env.Log

	mux := asynq.NewServeMux()
	mux.HandleFunc(constants.MarketDeploymentQueueName, env.ProcessDeploymentRequest)
	mux.HandleFunc(constants.TransactionMonitorQueueName, env.ProcessTransactionMonitorRequest)

	if err := env.QueueServer.Run(mux); err != nil {
		log.Error("could not run server: %v", err)
		os.Exit(1)
	}
}
