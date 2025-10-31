package models

import (
	"encoding/json"
	"time"
)

type MarketDeployRequest struct {
	MarketID string `json:"marketId"`
}

type TransactionMonitorRequest struct {
	MarketID       string `json:"marketId"`
	TransactionHex string `json:"transactionHex"`
}

type AdapterEventData struct {
	QuestionID        string
	ConditionID       string
	TokenIds          json.RawMessage
	DeployedTimestamp time.Time
}

type DeploymentStatusResponse struct {
	Status string `json:"status"`
}
