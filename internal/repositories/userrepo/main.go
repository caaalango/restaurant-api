package userrepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *UserRepository
	once     sync.Once
	table    string
)

type UserRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.User]
}

func New(conn *dbr.Connection) ports.UserRepository {
	table = "users"
	once.Do(func() {
		instance = &UserRepository{
			BaseRepository: baserepo.New[entities.User](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}

func (r *UserRepository) CreateWithCredencial(conf ports.CreateWithCredencialConf) (*entities.User, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	tx, err := session.Begin()
	if err != nil {
		return nil, err
	}

	var user *entities.User = conf.User
	_, err = tx.InsertInto(r.table).
		Record(user).
		Exec()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var credential *entities.Credential = conf.Credential
	credential.UserToken = user.Token
	_, err = tx.InsertInto("credentials").
		Columns("user_id", "token").
		Record(conf.Credential).
		Exec()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return user, nil
}
