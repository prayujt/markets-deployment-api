package handlers

import (
	"markets-api/internal/constants"
	"markets-api/internal/database"
)

func getMarketDeploymentStatus(market database.Market) constants.MarketDeploymentStatus {
	if market.PendingDeployment {
		return constants.MarketDeploymentStatusQueued
	} else if market.Deploying {
		return constants.MarketDeploymentStatusDeploying
	} else if market.DeployedTimestamp.Valid {
		return constants.MarketDeploymentStatusDeployed
	}
	// TODO: is this case possible?
	return constants.MarketDeploymentStatusQueued
}
