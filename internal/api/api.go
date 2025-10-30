package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"

	"markets-api/internal/constants"
	"markets-api/internal/database"
	"markets-api/internal/models"
	"markets-api/internal/services"
)

func (env *ClientEnvironment) GetDeploymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	questionID := vars["questionID"]
	status, err := services.GetKey[string](env.RedisClient, fmt.Sprintf(constants.MarketDeploymentStatusFormat, questionID))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid questionId: %v", err), http.StatusNotFound)
		return
	}

	res := models.DeploymentStatusResponse{
		QuestionID: questionID,
		Status:     status,
	}
	json.NewEncoder(w).Encode(res)
}

func (env *ClientEnvironment) QueueDeploymentRequest(w http.ResponseWriter, r *http.Request) {
	var req models.MarketDeployRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("could not decode reqeust body: %v", err), http.StatusBadRequest)
		return
	}
	// TODO: add cleaner validator
	if req.QuestionID == "" || req.ConditionID == "" || req.PositionIDYes == "" || req.PositionIDNo == "" {
		http.Error(w, "missing required fields in request body", http.StatusBadRequest)
		return
	}
	env.Log = env.Log.With("questionID", req.QuestionID, "conditionID", req.ConditionID)
	env.Log.Info("queuing deployment request", "questionID", req.QuestionID, "conditionID", req.ConditionID)

	// set separate status tracking key
	services.SetKey(env.RedisClient, fmt.Sprintf(constants.MarketDeploymentStatusFormat, req.QuestionID), constants.MarketDeploymentStatusQueued)

	// enqueue task
	services.PushToQueue(env.QueueClient, constants.MarketDeploymentQueueName, req)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"deploy triggered"}`))
}

func (env *ServerEnvironment) ProcessDeploymentRequest(ctx context.Context, t *asynq.Task) error {
	log := env.Log

	log.Info("processing deployment task", "payload", string(t.Payload()))

	var req models.MarketDeployRequest

	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return fmt.Errorf("could not decode task payload: %v", err)
	}

	env.Queries.CreateDeployment(context.Background(), database.CreateDeploymentParams{
		QuestionID:    req.QuestionID,
		ConditionID:   req.ConditionID,
		PositionIDYes: req.PositionIDYes,
		PositionIDNo:  req.PositionIDNo,
		Status:        string(constants.MarketDeploymentStatusQueued),
	})

	err := services.DeployMarket(log, env.ContractHTTP, env.ContractAuth, &req)
	if err != nil {
		log.Error("failed to deploy market", "error", err)
		return err
	}

	return nil
}
