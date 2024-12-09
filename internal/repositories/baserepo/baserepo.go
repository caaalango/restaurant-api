package baserepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo/dblogger"
	dbhdl "github.com/calango-productions/api/internal/repositories/baserepo/handlers"
	"github.com/gocraft/dbr/v2"
)

type BaseRepository[T any] struct {
	session *dbr.Session
	table   string
}

var (
	session *dbr.Session
	once    sync.Once
)

func New[T any](conn *dbr.Connection, table string) *BaseRepository[T] {
	once.Do(func() {
		session = conn.NewSession(&dblogger.LoggingEventReceiver{})
	})

	return &BaseRepository[T]{
		session: session,
		table:   table,
	}
}

func (r *BaseRepository[T]) Create(conf ports.CreateConf[T]) (*T, error) {
	err := dbhdl.New[T](r.session, r.table).Insert(&conf.Item)
	if err != nil {
		return nil, err
	}
	return &conf.Item, nil
}

func (r *BaseRepository[T]) CreateMany(conf ports.CreateManyConf[T]) error {
	err := dbhdl.New[T](r.session, r.table).InsertMany(conf.Items)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Exists(conf ports.ExistConf) (bool, error) {
	result, err := dbhdl.New[T](r.session, r.table).Get(conf.Key, conf.Value)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return true, nil
}

func (r *BaseRepository[T]) Get(conf ports.GetConf) (*T, error) {
	var entity T
	result, err := dbhdl.New[T](r.session, r.table).Get("Token", conf.Token)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("entity not found")
	}
	return &entity, nil
}

func (r *BaseRepository[T]) GetByKey(conf ports.GetByKeyConf) (*T, error) {
	var entity T
	result, err := dbhdl.New[T](r.session, r.table).Get(conf.Key, conf.Value)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("entity not found")
	}
	return &entity, nil
}

func (r *BaseRepository[T]) List(conf ports.ListConf) ([]T, error) {

	result, err := dbhdl.New[T](r.session, r.table).List(
		conf.HasPagination,
		conf.Size,
		conf.Page,
		conf.OnlyActives,
		conf.FilterToken,
		conf.FilterField,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepository[T]) Search(conf ports.SearchConf) ([]T, error) {
	result, err := dbhdl.New[T](r.session, r.table).Search(conf.Search, conf.Fields, conf.Size, conf.Page, conf.OnlyActives)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepository[T]) Update(conf ports.UpdateConf) error {
	err := dbhdl.New[T](r.session, r.table).Update("token", conf.Token, conf.Updates)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Inactivate(conf ports.InactivateConf) error {
	updates := map[string]interface{}{
		"active":     false,
		"updated_at": time.Now(),
	}
	err := dbhdl.New[T](r.session, r.table).Update("token", conf.Token, updates)
	if err != nil {
		return fmt.Errorf("failed to inactivate entity: %v", err)
	}
	return nil
}
