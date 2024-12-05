package entities

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Token     uuid.UUID `json:"token" db:"token"`
	Slug      string    `json:"slug" db:"slug"`
	Cnpj      string    `json:"cnpj" db:"cnpj"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
}
