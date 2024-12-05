package ports

import (
	"github.com/calango-productions/api/internal/core/entities"
)

type CreateWithCredencialConf struct {
	User       *entities.User
	Credential *entities.Credential
}

type UserRepository interface {
	BasePort[entities.User]
	CreateWithCredencial(CreateWithCredencialConf) (*entities.User, error)
}
