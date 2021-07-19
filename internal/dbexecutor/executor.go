package dbexecutor

import "database/sql"

type IDBExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}
