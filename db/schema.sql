CREATE TABLE markets (
    market_id TEXT PRIMARY KEY,
    question_id TEXT NULL,
    condition_id TEXT NULL,
	clob_token_ids JSONB NULL,
    pending_deployment BOOLEAN NOT NULL,
    deploying BOOLEAN NOT NULL,
    deploying_timestamp TIMESTAMP NULL,
    deployed_timestamp TIMESTAMP NULL
);
