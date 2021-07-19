package reporegistry

import "github.com/adalekin/otusk8s3/internal/repository/userrepo"

type InTransaction func(repoRegistry IRepoRegistry) (interface{}, error)

type IRepoRegistry interface {
	GetUserRepository() userrepo.IRepository
	DoInTransaction(txFunc InTransaction) (out interface{}, err error)
}
