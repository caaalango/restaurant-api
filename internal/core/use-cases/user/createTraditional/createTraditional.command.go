package createtradusercmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/pkg/encrypt"
	"github.com/calango-productions/api/pkg/regex"
)

func New(
	userRepository ports.UserRepository,
	crendencialRepository ports.CredentialRepository,
) *Command {
	return &Command{
		userRepository: userRepository,
	}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("create traditional user initiated: %v", params)

	isEmail := regex.IsValidEmail(params.Email)
	if !isEmail {
		log.Printf("field email is not email: %v", params.Email)
		return &Result{Success: false, Status: http.StatusBadRequest}, fmt.Errorf("field email is not email")
	}

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("error to get related entities: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("failed to create: %v", err)
	}

	if entities.User != nil {
		log.Printf("user already exists: %v", params)
		return &Result{Success: false, Status: http.StatusBadRequest}, fmt.Errorf("email already exists")
	}

	encrypted, err := encrypt.HashPassword(params.Password)
	if err != nil {
		log.Printf("error to encrypt password: %v", err)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("error to encrypt password")
	}

	buildedUser := c.buildUser(params)
	buildedCredential := c.buildCredential(encrypted)

	createdUser, err := c.persistEntities(buildedUser, buildedCredential)
	if err != nil {
		log.Printf("error to creation user: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("failed to create: %v", err)
	}

	log.Printf("create traditional user finished: %v", params)

	return &Result{Success: true, Status: http.StatusOK, Data: EntityIdentifier{UserToken: createdUser.Token.String()}}, nil
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	user, err := c.userRepository.GetByKey(ports.GetByKeyConf{
		Key:         "email",
		Value:       params.Email,
		OnlyActives: false,
	})
	if err != nil {
		return nil, err
	}

	return &RelatedEntities{User: user}, nil
}

func (c *Command) buildUser(params Params) *entities.User {
	return &entities.User{
		Email: params.Email,
	}
}

func (c *Command) buildCredential(encryptedPwd string) *entities.Credential {
	return &entities.Credential{
		Password: encryptedPwd,
		Type:     entities.CredentialType(entities.PASSWORD),
	}
}

func (c *Command) persistEntities(user *entities.User, credential *entities.Credential) (*entities.User, error) {
	user, err := c.userRepository.CreateWithCredencial(ports.CreateWithCredencialConf{User: user, Credential: credential})
	if err != nil {
		return nil, fmt.Errorf("failed to create: %v", err)
	}

	return user, nil
}
