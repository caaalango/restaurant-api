package entities

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	Token     uuid.UUID `json:"token" db:"token"`
	Name      string    `json:"name" db:"name"`
	Slogo     string    `json:"slogo" db:"slogo"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
}
