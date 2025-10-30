package api

import (
	"context"
	"net/http"

	"github.com/hibiken/asynq"

	"markets-api/internal/constants"
	"markets-api/internal/services"
)

func (h *ClientEnvironment) DeployHandler(w http.ResponseWriter, r *http.Request) {
	log := h.Log
	log.Info("deploy endpoint called")

	services.PushToQueue(h.QueueClient, constants.MarketDeploymentQueue, "hello")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"deploy triggered"}`))
}

func (h *ServerEnvironment) DeployHandler(ctx context.Context, t *asynq.Task) error {
	log := h.Log

	log.Info("processing deployment task", "payload", string(t.Payload()))

	return nil
}
