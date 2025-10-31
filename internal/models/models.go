package models

type MarketDeployRequest struct {
	MarketID string `json:"marketId"`
}

type TransactionMonitorRequest struct {
	MarketID       string `json:"marketId"`
	TransactionHex string `json:"transactionHex"`
}

type DeploymentStatusResponse struct {
	Status string `json:"status"`
}
