CREATE TABLE users (
    id UUID PRIMARY KEY,
    plan_id UUID NOT NULL REFERENCES plans(id),
    email VARCHAR(255) NOT NULL
);
