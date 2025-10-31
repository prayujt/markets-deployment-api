package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"

	"markets-api/internal/constants"
	"markets-api/internal/environment"
	"markets-api/internal/models"
	"markets-api/internal/services"
)

type ClientEnvironment struct {
	*environment.ClientEnvironment
}

type ServerEnvironment struct {
	*environment.ServerEnvironment
}

func NewClientEnvironment() *ClientEnvironment {
	return &ClientEnvironment{
		ClientEnvironment: environment.NewClientEnvironment(),
	}
}

func NewServerEnvironment() *ServerEnvironment {
	return &ServerEnvironment{
		ServerEnvironment: environment.NewServerEnvironment(),
	}
}

func (env *ClientEnvironment) GetDeploymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	marketID := vars["marketID"]
	market, err := env.Queries.GetMarketByID(context.Background(), marketID)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get market details: %v", err), http.StatusInternalServerError)
		return
	}

	res := models.DeploymentStatusResponse{
		Status: string(getMarketDeploymentStatus(market)),
	}
	json.NewEncoder(w).Encode(res)
}

func (env *ClientEnvironment) QueueDeploymentRequest(w http.ResponseWriter, r *http.Request) {
	var req models.MarketDeployRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("could not decode reqeust body: %v", err), http.StatusBadRequest)
		return
	}
	if req.MarketID == "" {
		http.Error(w, "missing required fields in request body", http.StatusBadRequest)
		return
	}
	env.LogWith("market_id", req.MarketID)
	env.Log.Info("queuing deployment request")

	// initialize deployment record
	env.Queries.SetMarketPendingDeployment(context.Background(), req.MarketID)

	// enqueue task
	services.PushToQueue(env.ClientEnvironment, constants.MarketDeploymentQueueName, req)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "deploy request queued"})
}

func (env *ServerEnvironment) ProcessDeploymentRequest(ctx context.Context, t *asynq.Task) error {
	log := env.Log
	log.Info("processing deployment task")

	var req models.MarketDeployRequest
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return fmt.Errorf("could not decode task payload: %v", err)
	}

	// update deployment status to deploying
	env.Queries.SetMarketDeploying(context.Background(), req.MarketID)

	err := services.DeployMarket(env.ServerEnvironment, &req)
	if err != nil {
		log.Error("failed to deploy market", "error", err)
		return err
	}

	return nil
}
