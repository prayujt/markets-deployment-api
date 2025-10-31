-- name: GetMarketByID :one
SELECT *
FROM markets
WHERE id = $1;

-- name: SetMarketPendingDeployment :exec
UPDATE markets
SET pending_deployment = TRUE
WHERE id = $1;

-- name: SetMarketDeploying :exec
UPDATE markets
SET deploying = TRUE,
	pending_deployment = FALSE,
    deploying_timestamp = NOW()
WHERE id = $1;

-- name: SetMarketDeployed :exec
UPDATE markets
SET deploying = FALSE,
    pending_deployment = FALSE,
    deployed_timestamp = NOW(),
    question_id = $2,
    condition_id = $3,
	clob_token_ids = $4
WHERE id = $1;
