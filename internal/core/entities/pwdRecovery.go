package entities

import (
	"time"

	"github.com/google/uuid"
)

type PasswordRecovery struct {
	Token     uuid.UUID `json:"token" db:"token"`
	UserToken uuid.UUID `json:"userToken" db:"user_token"`
	HashToken string    `json:"hashToken" db:"hash_token"`
	ExpiresAt time.Time `json:"expiresAt" db:"expires_at"`
	Active    bool      `json:"active" db:"active"`
}
