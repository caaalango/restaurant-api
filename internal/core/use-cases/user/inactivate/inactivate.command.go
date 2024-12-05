package inactivateusercmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/google/uuid"
)

type InactivateUserCommand struct {
	userRepository ports.UserRepository
}

func New(userRepository ports.UserRepository) *InactivateUserCommand {
	return &InactivateUserCommand{
		userRepository: userRepository,
	}
}

type InactivateUserParams struct {
	ActorToken string
}

type InactivateUserEntities struct {
	User *entities.User
}

type InactivateUserResult struct {
	Success bool
	Status  int
	Data    any
}

func (c *InactivateUserCommand) Execute(
	params InactivateUserParams,
) (*InactivateUserResult, error) {
	log.Printf("InactivateUserCommand initiated: %v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("CreateCandidateCommand failed: %v", params)
		return &InactivateUserResult{Success: false, Status: http.StatusInternalServerError}, err
	}
	if entities.User == nil {
		log.Printf("CreateCandidateCommand failed: %v", params)
		return &InactivateUserResult{Success: false, Status: http.StatusBadRequest}, fmt.Errorf("user not found")
	}

	err = c.guardAgainstUnauthorizedModifier(entities.User, params)
	if err != nil {
		log.Printf("GetCandidateCommand failed: %v", params)
		return &InactivateUserResult{Success: false, Status: http.StatusUnauthorized, Data: nil}, err
	}

	err = c.persistEntities(params)
	if err != nil {
		log.Printf("InactivateUserCommand failed: %v", params)
		return &InactivateUserResult{Success: false, Status: http.StatusUnauthorized, Data: nil}, err
	}

	log.Printf("InactivateUserCommand finished: %v", params)

	return &InactivateUserResult{Success: true}, nil
}

func (c *InactivateUserCommand) getRelatedEntities(params InactivateUserParams) (*InactivateUserEntities, error) {
	user, err := c.userRepository.Get(ports.GetConf{
		Token:       params.ActorToken,
		OnlyActives: true,
	})
	if err != nil {
		return nil, err
	}

	return &InactivateUserEntities{User: user}, nil
}

func (c *InactivateUserCommand) guardAgainstUnauthorizedModifier(user *entities.User, params InactivateUserParams) error {
	actorUuid, err := uuid.Parse(params.ActorToken)
	if err != nil {
		log.Printf("InactivateUserCommand failed: %v", params)
		return fmt.Errorf("failed to transform into uuid: %v", err)
	}

	if user.Token != actorUuid {
		return fmt.Errorf("unauthorized access")
	}

	return nil
}

func (c *InactivateUserCommand) persistEntities(params InactivateUserParams) error {
	err := c.userRepository.Inactivate(ports.InactivateConf{Token: params.ActorToken})
	if err != nil {
		log.Printf("InactivateUserCommand failed: %v", params)
		return err
	}

	return nil
}
