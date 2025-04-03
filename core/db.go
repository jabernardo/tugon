package core

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func GetDBInstance() *sql.DB {
	dbOnce.Do(func() {
		var (
			err    error
			exists bool
			driver string
			conn   string
		)

		if driver, exists = os.LookupEnv("SQL_DRIVER"); !exists {
			GetLoggerInstance().Error("[core.db] Environment `SQL_DRIVER` is not defined")
			os.Exit(1)
		}

		if conn, exists = os.LookupEnv("SQL_CONN"); !exists {
			GetLoggerInstance().Error("[core.db] Environment `SQL_DRIVER` is not defined")
			os.Exit(1)
		}

		db, err = sql.Open(driver, conn)

		if err != nil {
			GetLoggerInstance().Error("[core.db] Cannot open connection to database", "error", err)
			os.Exit(1)
		}
	})

	return db
}
