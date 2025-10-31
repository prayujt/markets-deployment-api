package models

type MarketDeployRequest struct {
	MarketID string `json:"marketId"`
}

type DeploymentStatusResponse struct {
	Status string `json:"status"`
}
