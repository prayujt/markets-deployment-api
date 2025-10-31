package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"

	"markets-api/internal/constants"
	"markets-api/internal/models"
	"markets-api/internal/services"
)

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

	marketExists, err := services.GetKey[bool](&env.ClientEnvironment.BaseEnvironment, fmt.Sprintf(constants.MarketRedisKeyFormat, req.MarketID))
	if err != nil {
		http.Error(w, fmt.Sprintf("could not check market existence: %v", err), http.StatusInternalServerError)
		return
	}
	if marketExists != nil && *marketExists {
		http.Error(w, "market is already pending deployment or deployed", http.StatusBadRequest)
		return
	}

	err = services.SetKey(&env.ClientEnvironment.BaseEnvironment, fmt.Sprintf(constants.MarketRedisKeyFormat, req.MarketID), true)
	if err != nil {
		env.Log.Warn("failed to set deployment status in redis", "error", err)
	}

	env.LogWith("market_id", req.MarketID)
	env.Log.Info("queuing deployment request")

	// initialize deployment record
	env.Queries.SetMarketPendingDeployment(context.Background(), req.MarketID)

	// enqueue task
	retryCount := 3
	services.PushToQueue(&env.ClientEnvironment.BaseEnvironment, constants.MarketDeploymentQueueName, req, &retryCount)

	json.NewEncoder(w).Encode(map[string]string{"message": "market deployment queued"})
}

func (env *ServerEnvironment) ProcessDeploymentRequest(ctx context.Context, t *asynq.Task) error {
	log := env.Log
	log.Info("processing deployment task")

	var req models.MarketDeployRequest
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return fmt.Errorf("could not decode deployment request task payload: %v", err)
	}

	// update deployment status to deploying
	err := env.Queries.SetMarketDeploying(context.Background(), req.MarketID)
	if err != nil {
		log.Error("failed to update market status to deploying", "error", err)
		return err
	}

	tx, err := services.RunAdapterInitialize(env.ServerEnvironment, &req)
	if err != nil {
		log.Error("failed to deploy market", "error", err)
		return err
	}

	services.PushToQueue(&env.ServerEnvironment.BaseEnvironment, constants.TransactionMonitorQueueName, models.TransactionMonitorRequest{
		MarketID:       req.MarketID,
		TransactionHex: tx,
	}, nil)

	return nil
}

func (env *ServerEnvironment) ProcessTransactionMonitorRequest(ctx context.Context, t *asynq.Task) error {
	log := env.Log
	log.Info("processing deployment task")

	var req models.TransactionMonitorRequest
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return fmt.Errorf("could not decode transaction monitor task payload: %v", err)
	}

	receipt, err := services.BlockUntilTransactionMined(env.ServerEnvironment, req.TransactionHex)
	if err != nil {
		log.Error("failed to monitor transaction", "error", err)
		return err
	}

	data, err := services.ParseInitializeReceipt(env.ServerEnvironment, receipt)
	if err != nil {
		log.Error("failed to parse initialize receipt", "error", err)
		return err
	}

	// update market record
	err = services.CompleteMarketDeployment(env.ServerEnvironment, req.MarketID, data)
	if err != nil {
		log.Error("failed to update records in database", "error", err)
		return err
	}

	env.Log.Info("market deployed successfully", "market_id", req.MarketID)

	return nil
}
