package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	GUEST  UserRole = "guest"
	CLIENT UserRole = "client"
	WAITER UserRole = "waiter"
	CHEF   UserRole = "chef"
	ADMIN  UserRole = "admin"
)

type User struct {
	Token       uuid.UUID `json:"token" db:"token"`
	ClientToken uuid.UUID `json:"clientToken" db:"client_token"`
	Email       string    `json:"email" db:"email"`
	Role        UserRole  `json:"role" db:"role"`
	Permissions string    `json:"permissions" db:"permissions"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
	LastLogin   time.Time `json:"lastLogin" db:"last_login"`
	Active      bool      `json:"active" db:"active"`
}
