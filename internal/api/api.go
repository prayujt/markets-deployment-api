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

	questionID := vars["questionID"]
	status, err := env.Queries.GetDeploymentStatus(context.Background(), questionID)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get deployment status: %v", err), http.StatusInternalServerError)
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

	// initialize deployment record
	env.Queries.CreateDeployment(context.Background(), database.CreateDeploymentParams{
		QuestionID:    req.QuestionID,
		ConditionID:   req.ConditionID,
		PositionIDYes: req.PositionIDYes,
		PositionIDNo:  req.PositionIDNo,
		Status:        string(constants.MarketDeploymentStatusQueued),
	})

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
	// add to set to track active deployments from this instance
	env.QuestionIDSet.Add(req.QuestionID)

	// update deployment status to deploying
	env.Queries.UpdateDeploymentStatus(context.Background(), database.UpdateDeploymentStatusParams{
		QuestionID: req.QuestionID,
		Status:     string(constants.MarketDeploymentStatusDeploying),
	})

	err := services.DeployMarket(env.ServerEnvironment, &req)
	if err != nil {
		log.Error("failed to deploy market", "error", err)
		return err
	}

	return nil
}
