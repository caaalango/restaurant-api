package entities

import (
	"time"

	"github.com/google/uuid"
)

type Dish struct {
	Token        uuid.UUID `json:"token" db:"token"`
	ClientToken  uuid.UUID `json:"clientToken" db:"client_token"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Restrictions string    `json:"restrictions" db:"restrictions"`
	Price        int       `json:"price" db:"price"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
	Active       bool      `json:"active" db:"active"`
}
