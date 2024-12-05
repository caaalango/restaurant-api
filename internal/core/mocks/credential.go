package mocks

import (
	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/pkg/random"
	"github.com/stretchr/testify/mock"
)

type CredentialRepository struct {
	mock.Mock
}

func (m *CredentialRepository) CreateWithCredencial(conf ports.CreateWithCredencialConf) (*entities.Credential, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.Credential), args.Error(1)
}

func (m *CredentialRepository) Exists(conf ports.ExistConf) (bool, error) {
	args := m.Called(conf)
	return args.Bool(0), args.Error(1)
}

func (m *CredentialRepository) Create(conf ports.CreateConf[entities.Credential]) (*entities.Credential, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.Credential), args.Error(1)
}

func (m *CredentialRepository) CreateMany(conf ports.CreateManyConf[entities.Credential]) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *CredentialRepository) Get(conf ports.GetConf) (*entities.Credential, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.Credential), args.Error(1)
}

func (m *CredentialRepository) GetByKey(conf ports.GetByKeyConf) (*entities.Credential, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.Credential), args.Error(1)
}

func (m *CredentialRepository) List(conf ports.ListConf) ([]entities.Credential, error) {
	args := m.Called(conf)
	return args.Get(0).([]entities.Credential), args.Error(1)
}

func (m *CredentialRepository) Search(conf ports.SearchConf) ([]entities.Credential, error) {
	args := m.Called(conf)
	return args.Get(0).([]entities.Credential), args.Error(1)
}

func (m *CredentialRepository) Update(conf ports.UpdateConf) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *CredentialRepository) Inactivate(conf ports.InactivateConf) error {
	args := m.Called(conf)
	return args.Error(0)
}

func GenerateCredenditialRepo() *CredentialRepository {
	userRepo := new(CredentialRepository)
	userRepo.On("GetByKey", mock.Anything).Return(GenerateCredential(), nil)
	userRepo.On("Exists", mock.Anything).Return(true, nil)
	userRepo.On("CreateWithCredencial", mock.Anything).Return(GenerateCredential(), nil)
	return userRepo
}

func GenerateCredential() *entities.Credential {
	return &entities.Credential{
		Password: random.String(12),
		Type:     entities.PASSWORD,
	}
}
