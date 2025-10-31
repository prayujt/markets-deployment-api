package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sqlc-dev/pqtype"

	"markets-api/internal/constants"
	"markets-api/internal/database"
	"markets-api/internal/environment"
	"markets-api/internal/models"
)

func RunAdapterInitialize(env *environment.ServerEnvironment, market *models.MarketDeployRequest) (string, error) {
	marketData, err := env.Queries.GetMarketByID(context.Background(), market.MarketID)
	if err != nil {
		env.Log.Error("failed to fetch market from database", "error", err)
		return "", err
	}

	required := big.NewInt(constants.AdapterRewardSize)

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	allow, err := env.USDC.Allowance(&bind.CallOpts{Context: ctx}, env.ContractAuth.From, common.HexToAddress(constants.AdapterAddressHex))
	if err != nil {
		env.Log.Error("allowance check failed", "err", err)
		return "", err
	}

	if allow.Cmp(required) < 0 {
		env.Log.Info("current allowance insufficient", "current", allow.String(), "required", required.String())
		env.Log.Info("approving USDC for adapter", "spender", constants.AdapterAddressHex, "amount", required.String())

		txApprove, err := env.USDC.Approve(env.ContractAuth, common.HexToAddress(constants.AdapterAddressHex), required)
		if err != nil {
			env.Log.Error("approve failed", "err", err)
			return "", err
		}
		env.Log.Info("approve tx sent", "tx", txApprove.Hash().Hex())

		if _, err := bind.WaitMined(context.Background(), env.EthHTTP, txApprove); err != nil {
			env.Log.Error("approve tx failed to mine", "err", err)
			return "", err
		}
	}

	// TODO: construct data properly
	ancillaryData := []byte(marketData.Question + "|" + marketData.Description + "|" + string(marketData.Outcomes))
	tx, err := env.ContractHTTP.Initialize(
		env.ContractAuth,
		ancillaryData,
		constants.USDCAddress,
		big.NewInt(constants.AdapterRewardSize),
		big.NewInt(constants.AdapterProposalBond),
		big.NewInt(constants.AdapterLiveness))
	if err != nil {
		env.Log.Error("failed to deploy market contract", "error", err)
		return "", err
	}
	hex := tx.Hash().Hex()
	env.Log.Info("adapter initialize transaction sent", "tx", hex)
	return hex, nil
}

func BlockUntilTransactionMined(env *environment.ServerEnvironment, hex string) error {
	hash := common.HexToHash(hex)
	var receipt *types.Receipt

	// max 10 minutes wait, otherwise will requeue
	outerCtx, outerCancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer outerCancel()

	for {
		select {
		case <-outerCtx.Done():
			return fmt.Errorf("timed out waiting for transaction to be mined: %w", outerCtx.Err())
		default:
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		r, err := env.EthHTTP.TransactionReceipt(ctx, hash)
		cancel()
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(2 * time.Second)
				continue
			}
			return fmt.Errorf("failed to get transaction receipt: %v", err)
		}
		receipt = r
		break
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction failed with status: %v", receipt.Status)
	}

	return nil
}

// func SubscribeMarketDeployments(env *environment.ServerEnvironment, events chan *contracts.ContractsMarketDeployed) (event.Subscription, error) {
// 	sub, err := env.ContractWS.WatchMarketDeployed(&bind.WatchOpts{Context: context.Background()}, events, nil)
// 	if err != nil {
// 		env.Log.Error("failed to subscribe to MarketDeployed events", "error", err)
// 		return nil, err
// 	}
// 	return sub, nil
// }

func CompleteMarketDeployment(env *environment.ServerEnvironment, marketID string) {
	err := env.Queries.SetMarketDeployed(context.Background(), database.SetMarketDeployedParams{
		ID:          marketID,
		QuestionID:  sql.NullString{Valid: true, String: "sample_question_id"},
		ConditionID: sql.NullString{Valid: true, String: "sample_condition_id"},
		ClobTokenIds: pqtype.NullRawMessage{
			Valid:      true,
			RawMessage: []byte(`["token1","token2"]`),
		},
	})
	if err != nil {
		env.Log.Error("failed to update deployment status in database", "error", err)
		return
	}
	env.Log.Info("market deployment completed", "questionID", marketID)
}
