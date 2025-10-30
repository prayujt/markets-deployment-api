package constants

const MarketDeploymentQueueName = "market_deployment_queue"

const MarketDeploymentStatusFormat = "deployment_status:%s"

type MarketDeploymentStatus string

const (
	MarketDeploymentStatusQueued    MarketDeploymentStatus = "QUEUED"
	MarketDeploymentStatusDeploying MarketDeploymentStatus = "DEPLOYING"
	MarketDeploymentStatusDeployed  MarketDeploymentStatus = "DEPLOYED"
)
