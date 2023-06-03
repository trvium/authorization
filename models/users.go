package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID     uuid.UUID `json:"id"`
	PlanID uuid.UUID `json:"plan_id"`
	Email  string    `json:"email"`
}
