package main

import (
	"os"

	"github.com/hibiken/asynq"

	"markets-api/internal/api"
	"markets-api/internal/constants"
)

func main() {
	env := api.NewServerEnvironment()
	log := env.Log

	mux := asynq.NewServeMux()
	mux.HandleFunc(constants.MarketDeploymentQueueName, env.ProcessDeploymentRequest)

	if err := env.QueueServer.Run(mux); err != nil {
		log.Error("could not run server: %v", err)
		os.Exit(1)
	}
}
