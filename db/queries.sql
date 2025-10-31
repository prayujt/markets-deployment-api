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
    deployed_timestamp = $2,
    question_id = $3,
    condition_id = $4,
	clob_token_ids = $5
WHERE id = $1;
