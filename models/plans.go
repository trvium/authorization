package models

import (
	"github.com/google/uuid"
)

type Plan struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Quota int       `json:"quota"`
}
