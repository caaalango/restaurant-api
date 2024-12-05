package dbhdl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/calango-productions/api/internal/constants"
	"github.com/gocraft/dbr/v2"
)

type CoreHandler[T any] struct {
	session *dbr.Session
	table   string
}

func New[T any](db *dbr.Session, table string) *CoreHandler[T] {
	return &CoreHandler[T]{session: db, table: table}
}

func (h *CoreHandler[T]) Insert(record *T) error {
	tx, err := h.session.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.InsertInto(h.table).Columns(getColumns(record)...).Record(record).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *CoreHandler[T]) InsertMany(records []T) error {
	tx, err := h.session.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	chunks := splitIntoChunks(records, constants.SIZE_OF_INSERT_MANY_ENTITIES)

	for _, chunk := range chunks {
		_, err := tx.InsertInto(h.table).Columns(getColumns(chunk)...).Record(chunk).Exec()
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (h *CoreHandler[T]) List(size int, page int, onlyActives bool, clientToken string) ([]T, error) {
	var results []T

	offset := (page - 1) * size

	query := h.session.Select("*").
		From(h.table).
		Limit(uint64(size)).
		Offset(uint64(offset))

	if onlyActives {
		query = query.Where("active = ?", true)
	}
	if clientToken != "" {
		query = query.Where("client_token = ?", clientToken)
	}

	_, err := query.Load(&results)
	if err != nil {
		return nil, fmt.Errorf("failed to list records from table %s: %v", h.table, err)
	}

	return results, nil
}

func (h *CoreHandler[T]) Search(search string, fields []string, size int, page int, onlyActives bool) ([]T, error) {
	var results []T

	offset := (page - 1) * size

	query := h.session.Select("*").
		From(h.table).
		Limit(uint64(size)).
		Offset(uint64(offset))

	if onlyActives {
		query = query.Where("active = ?", true)
	}

	if search != "" && len(fields) > 0 {
		var conditions []string
		var args []interface{}

		for _, field := range fields {
			conditions = append(conditions, fmt.Sprintf("%s LIKE ?", field))
			args = append(args, "%"+search+"%")
		}

		query = query.Where(strings.Join(conditions, " OR "), args...)
	}

	_, err := query.Load(&results)
	if err != nil {
		return nil, fmt.Errorf("failed to list records from table %s: %v", h.table, err)
	}

	return results, nil
}

func (h *CoreHandler[T]) Get(field string, value interface{}) (*T, error) {
	var result T
	_, err := h.session.Select("*").
		From(h.table).
		Where(fmt.Sprintf("%s = ?", field), value).
		Load(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (h *CoreHandler[T]) Update(field string, value interface{}, updates map[string]interface{}) error {
	tx, err := h.session.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update(h.table).
		SetMap(updates).
		Where(dbr.Eq(field, value)).
		Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *CoreHandler[T]) Delete(field string, value interface{}) error {
	tx, err := h.session.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.DeleteFrom(h.table).
		Where(dbr.Eq(field, value)).
		Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *CoreHandler[T]) Inactivate(field string, value interface{}) error {
	tx, err := h.session.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update(h.table).
		Set("active", false).
		Where(dbr.Eq(field, value)).
		Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func getColumns(record interface{}) []string {
	val := reflect.ValueOf(record).Elem()
	var columns []string

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			columns = append(columns, dbTag)
		}
	}
	return columns
}

func splitIntoChunks[T any](records []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(records); i += chunkSize {
		end := i + chunkSize
		if end > len(records) {
			end = len(records)
		}
		chunks = append(chunks, records[i:end])
	}
	return chunks
}
