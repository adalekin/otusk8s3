package reporegistry

import (
	"database/sql"

	"github.com/adalekin/otusk8s3/internal/dbexecutor"
	"github.com/adalekin/otusk8s3/internal/repository/userrepo"
	"github.com/pkg/errors"
)

type pgRepositoryRegistry struct {
	db         *sql.DB
	dbexecutor dbexecutor.IDBExecutor
}

func NewPostgreSQL(db *sql.DB) IRepoRegistry {
	return pgRepositoryRegistry{
		db: db,
	}
}

func (r pgRepositoryRegistry) GetUserRepository() userrepo.IRepository {
	if r.dbexecutor != nil {
		return userrepo.NewPostgreSQL(r.dbexecutor)
	}

	return userrepo.NewPostgreSQL(r.db)
}

func (r pgRepositoryRegistry) DoInTransaction(txFunc InTransaction) (out interface{}, err error) {
	var tx *sql.Tx
	registry := r
	if r.dbexecutor == nil {
		tx, err = r.db.Begin()
		if err != nil {
			return
		}
		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p) // re-throw panic after Rollback
			} else if err != nil {
				xerr := tx.Rollback() // err is non-nil; don't change it
				if xerr != nil {
					err = errors.Wrap(err, xerr.Error())
				}
			} else {
				err = tx.Commit() // err is nil; if Commit returns error update err
			}
		}()
		registry = pgRepositoryRegistry{
			db:         r.db,
			dbexecutor: tx,
		}
	}
	out, err = txFunc(registry)
	return
}
