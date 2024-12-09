package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Token           uuid.UUID `json:"commentToken" db:"comment_token"`
	RestaurantToken uuid.UUID `json:"restaurantToken" db:"restarant_token"`
	DishToken       uuid.UUID `json:"dishToken" db:"dish_token"`
	ReviewToken     uuid.UUID `json:"reviewToken" db:"review_token"`
	UserToken       uuid.UUID `json:"userToken" db:"user_token"`
	Message         string    `json:"message" db:"message"`
	Note            int       `json:"note" db:"note"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
	Active          bool      `json:"active" db:"active"`
}
