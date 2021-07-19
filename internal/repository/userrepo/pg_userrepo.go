package userrepo

import (
	"context"

	"github.com/adalekin/otusk8s3/internal/dbexecutor"
	"github.com/adalekin/otusk8s3/internal/domain"
)

type pgRepository struct {
	db dbexecutor.DBExecutor
}

func NewPostgreSQL(db dbexecutor.DBExecutor) IRepository {
	return pgRepository{
		db: db,
	}
}

func (r pgRepository) Create(ctx context.Context, req CreateRequest) (res *domain.User, err error) {
	query := "insert into users (uuid, name) values (?, ?);"
	res = &domain.Person{UUID: uuid.NewUUID().String(), Name: req.Name}

	_, err = r.db.Exec(query, res.UUID, res.Name)

	if err != nil {
		return nil, err
	}

	return
}

func (r pgRepository) FindAll(ctx context.Context) (res []*domain.User, err error) {
	return
}

func (r pgRepository) FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error) {
	return
}

func (r pgRepository) Delete(ctx context.Context, req DeleteRequest) (err error) {
	return
}
