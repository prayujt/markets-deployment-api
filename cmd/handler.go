package main

import (
	"os"

	"github.com/hibiken/asynq"

	"markets-api/internal/constants"
	"markets-api/internal/handlers"
	"markets-api/internal/services"
)

func main() {
	env := handlers.NewServerEnvironment()
	log := env.Log

	mux := asynq.NewServeMux()
	mux.HandleFunc(constants.MarketDeploymentQueueName, env.ProcessDeploymentRequest)
	mux.HandleFunc(constants.TransactionMonitorQueueName, env.ProcessTransactionMonitorRequest)

	err := services.ResetAllowance(env.ServerEnvironment)
	if err != nil {
		log.Error("could not reset allowance: %v", err)
		os.Exit(1)
	}

	if err := env.QueueServer.Run(mux); err != nil {
		log.Error("could not run server: %v", err)
		os.Exit(1)
	}
}
