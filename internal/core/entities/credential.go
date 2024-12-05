package entities

import (
	"time"

	"github.com/google/uuid"
)

type CredentialType string

const (
	PASSWORD  = "password"
	GOOGLE    = "google"
	MICROSOFT = "microsoft"
	LINKEDIN  = "linkedin"
)

type Credential struct {
	Token     uuid.UUID      `json:"token" db:"token"`
	UserToken uuid.UUID      `json:"userToken" db:"user_token"`
	Password  string         `json:"password" db:"password"`
	Type      CredentialType `json:"type" db:"type"`
	SocialID  string         `json:"socialId" db:"social_id"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" db:"updated_at"`
	Active    bool           `json:"active" db:"active"`
}
