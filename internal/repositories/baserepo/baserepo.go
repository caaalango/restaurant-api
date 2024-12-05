package baserepo

import (
	"fmt"
	"time"

	"github.com/calango-productions/api/internal/core/ports"
	dbhdl "github.com/calango-productions/api/internal/repositories/baserepo/handlers"
	"github.com/gocraft/dbr/v2"
)

type BaseRepository[T any] struct {
	conn  *dbr.Connection
	table string
}

func New[T any](conn *dbr.Connection, table string) *BaseRepository[T] {
	return &BaseRepository[T]{
		conn:  conn,
		table: table,
	}
}

func (r *BaseRepository[T]) Create(conf ports.CreateConf[T]) (*T, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	err := dbhdl.New[T](session, r.table).Insert(&conf.Item)
	if err != nil {
		return nil, err
	}
	return &conf.Item, nil
}

func (r *BaseRepository[T]) CreateMany(conf ports.CreateManyConf[T]) error {
	session := r.conn.NewSession(nil)
	defer session.Close()

	err := dbhdl.New[T](session, r.table).InsertMany(conf.Items)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Exists(conf ports.ExistConf) (bool, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	result, err := dbhdl.New[T](session, r.table).Get(conf.Key, conf.Value)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return true, nil
}

func (r *BaseRepository[T]) Get(conf ports.GetConf) (*T, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	var application T
	result, err := dbhdl.New[T](session, r.table).Get("Token", conf.Token)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("application not found")
	}
	return &application, nil
}

func (r *BaseRepository[T]) GetByKey(conf ports.GetByKeyConf) (*T, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	var application T
	result, err := dbhdl.New[T](session, r.table).Get(conf.Key, conf.Value)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("application not found")
	}
	return &application, nil
}

func (r *BaseRepository[T]) List(conf ports.ListConf) ([]T, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	result, err := dbhdl.New[T](session, r.table).List(conf.Size, conf.Page, conf.OnlyActives, conf.ClientToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepository[T]) Search(conf ports.SearchConf) ([]T, error) {
	session := r.conn.NewSession(nil)
	defer session.Close()

	result, err := dbhdl.New[T](session, r.table).Search(conf.Search, conf.Fields, conf.Size, conf.Page, conf.OnlyActives)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepository[T]) Update(conf ports.UpdateConf) error {
	session := r.conn.NewSession(nil)
	defer session.Close()

	err := dbhdl.New[T](session, r.table).Update("token", conf.Token, conf.Updates)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Inactivate(conf ports.InactivateConf) error {
	session := r.conn.NewSession(nil)
	defer session.Close()

	updates := map[string]interface{}{
		"active":     false,
		"updated_at": time.Now(),
	}
	err := dbhdl.New[T](session, r.table).Update("token", conf.Token, updates)
	if err != nil {
		return fmt.Errorf("failed to inactivate application: %v", err)
	}
	return nil
}
