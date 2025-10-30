CREATE TABLE deployments (
    question_id TEXT PRIMARY KEY,
    condition_id TEXT NOT NULL,
    position_id_yes TEXT NOT NULL,
    position_id_no TEXT NOT NULL,
    status TEXT NOT NULL
);
