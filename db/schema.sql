CREATE TABLE markets (
    id TEXT PRIMARY KEY,
    question TEXT NOT NULL,
    description TEXT NOT NULL,
	outcomes JSONB NOT NULL,
    uma_bond TEXT NOT NULL,
    uma_reward TEXT NOT NULL,
    question_id TEXT NULL,
    condition_id TEXT NULL,
	clob_token_ids JSONB NULL,
    pending_deployment BOOLEAN NOT NULL,
    deploying BOOLEAN NOT NULL,
    deploying_timestamp TIMESTAMP NULL,
    deployed_timestamp TIMESTAMP NULL
);
