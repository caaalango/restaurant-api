package entities

import (
	"time"

	"github.com/calango-productions/api/internal/core/enums"
	"github.com/google/uuid"
)

type User struct {
	Token     uuid.UUID      `json:"token" db:"token"`
	Email     string         `json:"email" db:"email"`
	Icon      enums.UserIcon `json:"icon" db:"icon"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" db:"updated_at"`
	LastLogin time.Time      `json:"lastLogin" db:"last_login"`
	Active    bool           `json:"active" db:"active"`
}
