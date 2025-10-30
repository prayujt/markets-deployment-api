-- name: CreateDeployment :exec
INSERT INTO deployments (question_id, condition_id, position_id_yes, position_id_no, status)
VALUES ($1, $2, $3, $4, $5);

-- name: GetDeploymentStatus :one
SELECT status
FROM deployments
WHERE question_id = $1;

-- name: UpdateDeploymentStatus :exec
UPDATE deployments
SET status = $2
WHERE question_id = $1;
