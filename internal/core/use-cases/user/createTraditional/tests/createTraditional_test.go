package createtradusertest

import (
	"errors"
	"net/http"
	"testing"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/mocks"
	"github.com/calango-productions/api/internal/core/ports"
	createtradusercmd "github.com/calango-productions/api/internal/core/use-cases/user/createTraditional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildGenericDependecies() (*createtradusercmd.Command, *mocks.UserRepository, *mocks.CredentialRepository) {
	userRepo := new(mocks.UserRepository)
	credRepo := new(mocks.CredentialRepository)

	cmd := createtradusercmd.New(userRepo, credRepo)

	return cmd, userRepo, credRepo
}

func buildGenericParams() createtradusercmd.Params {
	return createtradusercmd.Params{
		Email:    "test@example.com",
		Password: "password123",
	}
}

func TestSuccessfulCreation(t *testing.T) {
	cmd, userRepo, credRepo := buildGenericDependecies()

	params := buildGenericParams()

	userRepo.
		On("GetByKey", ports.GetByKeyConf{
			Key:         "email",
			Value:       params.Email,
			OnlyActives: false,
		}).
		Return(nil, nil)

	userRepo.
		On("CreateWithCredencial", mock.Anything).
		Return(&entities.User{Email: params.Email}, nil)

	credRepo.
		On("Create", mock.Anything).
		Return(&entities.Credential{Password: "encryptedPassword"}, nil)

	result, err := cmd.Execute(params)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, http.StatusOK, result.Status)
}

func TestUserEmailAlreadyExists(t *testing.T) {
	cmd, userRepo, _ := buildGenericDependecies()

	params := buildGenericParams()

	userRepo.On("GetByKey", ports.GetByKeyConf{
		Key:         "email",
		Value:       params.Email,
		OnlyActives: false,
	}).Return(&entities.User{Email: params.Email}, nil)

	result, err := cmd.Execute(params)

	assert.Error(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, http.StatusBadRequest, result.Status)
}

func TestUserCreateFail(t *testing.T) {
	cmd, userRepo, _ := buildGenericDependecies()

	params := buildGenericParams()

	userRepo.
		On("GetByKey", ports.GetByKeyConf{
			Key:         "email",
			Value:       params.Email,
			OnlyActives: false,
		}).
		Return(nil, nil)

	userRepo.
		On("CreateWithCredencial", mock.Anything).
		Return(nil, errors.New("failed to create"))

	result, err := cmd.Execute(params)

	assert.Error(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, http.StatusInternalServerError, result.Status)
}

func TestCredentialCreateFail(t *testing.T) {
	cmd, userRepo, credRepo := buildGenericDependecies()

	params := buildGenericParams()

	userRepo.
		On("GetByKey", ports.GetByKeyConf{
			Key:         "email",
			Value:       params.Email,
			OnlyActives: false,
		}).
		Return(nil, nil)

	userRepo.
		On("CreateWithCredencial", mock.Anything).
		Return(&entities.User{Email: params.Email}, nil)

	credRepo.
		On("Create", mock.Anything).
		Return(nil, errors.New("failed to create"))

	result, err := cmd.Execute(params)

	assert.Error(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, http.StatusInternalServerError, result.Status)
}
