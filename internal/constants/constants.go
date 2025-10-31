package constants

import "github.com/ethereum/go-ethereum/common"

const MarketDeploymentQueueName = "market_deployment_queue"
const TransactionMonitorQueueName = "transaction_monitor_queue"

type MarketDeploymentStatus string

const (
	MarketDeploymentStatusQueued    MarketDeploymentStatus = "QUEUED"
	MarketDeploymentStatusDeploying MarketDeploymentStatus = "DEPLOYING"
	MarketDeploymentStatusDeployed  MarketDeploymentStatus = "DEPLOYED"
)

const AdapterAddressHex = "0x65070BE91477460D8A7AeEb94ef92fe056C2f2A7"

// Amoy
var USDCAddress = common.HexToAddress("0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582")
var CTFAddress = common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB")

// Mainnet
// var USDCAddress = common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")
// var CTFAddress = common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045")

/**
 * TODO: check AdapterRewardSize value
 * is it constant or should it be passed in the request?
 */

const AdapterRewardSize = 1
const AdapterProposalBond = 500
const AdapterLiveness = 7200
