package models

import (
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
	DeployedTimestamp time.Time
}

type DeploymentStatusResponse struct {
	Status string `json:"status"`
}
