CREATE TABLE api_keys (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    key VARCHAR(64) NOT NULL,
    valid BOOLEAN NOT NULL,
    quota_used INT NOT NULL
);
