package ports

import "github.com/calango-productions/api/internal/core/entities"

type ClientRepository interface {
	BasePort[entities.Client]
}
