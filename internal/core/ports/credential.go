package ports

import "github.com/calango-productions/api/internal/core/entities"

type CredentialRepository interface {
	BasePort[entities.Credential]
}
