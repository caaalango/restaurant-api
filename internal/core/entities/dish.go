package entities

import (
	"database/sql"
	"time"

	"github.com/calango-productions/api/internal/core/enums"
	"github.com/google/uuid"
)

type Dish struct {
	Token            uuid.UUID          `json:"token" db:"token"`
	RestaurantToken  uuid.UUID          `json:"restaurantToken" db:"restaurant_token"`
	Category         enums.DishCategory `json:"category" db:"category"`
	Title            string             `json:"title" db:"title"`
	QuickDescription string             `json:"quickDescription" db:"quick_description"`
	LongDescription  string             `json:"longDescription" db:"long_description"`
	Restrictions     sql.NullString     `json:"restrictions" db:"restrictions"`
	Price            string             `json:"price" db:"price"`
	ImageUrl         string             `json:"imageUrl" db:"image_url"`
	VideoUrl         sql.NullString     `json:"videoUrl" db:"video_url"`
	CreatedAt        time.Time          `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time          `json:"updatedAt" db:"updated_at"`
	Active           bool               `json:"active" db:"active"`
}
