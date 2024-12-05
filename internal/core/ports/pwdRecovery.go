package ports

import (
	"github.com/calango-productions/api/internal/core/entities"
)

type PasswordRecoveryRepository interface {
	BasePort[entities.PasswordRecovery]
}
