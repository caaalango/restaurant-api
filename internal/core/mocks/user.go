package mocks

import (
	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) CreateWithCredencial(conf ports.CreateWithCredencialConf) (*entities.User, error) {
	args := m.Called(conf)
	var user *entities.User
	if u := args.Get(0); u != nil {
		user = u.(*entities.User)
	}
	return user, args.Error(1)
}

func (m *UserRepository) Exists(conf ports.ExistConf) (bool, error) {
	args := m.Called(conf)
	return args.Bool(0), args.Error(1)
}

func (m *UserRepository) Create(conf ports.CreateConf[entities.User]) (*entities.User, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *UserRepository) CreateMany(conf ports.CreateManyConf[entities.User]) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *UserRepository) Get(conf ports.GetConf) (*entities.User, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *UserRepository) GetByKey(conf ports.GetByKeyConf) (*entities.User, error) {
	args := m.Called(conf)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *UserRepository) List(conf ports.ListConf) ([]entities.User, error) {
	args := m.Called(conf)
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *UserRepository) Search(conf ports.SearchConf) ([]entities.User, error) {
	args := m.Called(conf)
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *UserRepository) Update(conf ports.UpdateConf) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *UserRepository) Inactivate(conf ports.InactivateConf) error {
	args := m.Called(conf)
	return args.Error(0)
}
