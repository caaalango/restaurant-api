package coreconn

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

var CORE_MAX_CONN = 50
var CORE_MAX_INACTIVATE_CONN = 20
var CORE_MAX_TIME_TO_REUSE_CONN = 20 * time.Minute
var CORE_MAX_TIME_TO_INATIVE_CONN = 10 * time.Minute

func ConnectCoreDatabase(dsn string) (*dbr.Connection, error) {
	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open database connection: %v", err)
	}

	conn.SetMaxOpenConns(CORE_MAX_CONN)
	conn.SetMaxIdleConns(CORE_MAX_INACTIVATE_CONN)
	conn.SetConnMaxLifetime(CORE_MAX_TIME_TO_REUSE_CONN)
	conn.SetConnMaxIdleTime(CORE_MAX_TIME_TO_INATIVE_CONN)

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	fmt.Println("Postgres connection established.")

	return conn, nil
}
