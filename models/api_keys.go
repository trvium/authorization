package models

import (
	"github.com/google/uuid"
)

type ApiKey struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Key       string    `json:"key"`
	Valid     bool      `json:"validity"`
	QuotaUsed int       `json:"quota_used"`
}
