package models

import (
	"github.com/google/uuid"
)

type Plan struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name  string    `json:"name"`
	Limit int       `json:"limit"`
}
