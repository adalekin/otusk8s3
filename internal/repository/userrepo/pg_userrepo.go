package userrepo

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/adalekin/otusk8s3/internal/dbexecutor"
	"github.com/adalekin/otusk8s3/internal/domain"
)

type pgRepository struct {
	db dbexecutor.IDBExecutor
}

func NewPostgreSQL(db dbexecutor.IDBExecutor) IRepository {
	return pgRepository{
		db: db,
	}
}

func (r pgRepository) Create(ctx context.Context, req CreateRequest) (res *domain.User, err error) {
	query := "INSERT INTO users (name) VALUES (?);"
	res = &domain.User{Name: req.Name}

	_, err = r.db.Exec(query, res.Name)

	if err != nil {
		return nil, err
	}

	return
}

func (r pgRepository) FindAll(ctx context.Context) (res []*domain.User, err error) {
	query := "SELECT * FROM users;"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user domain.User

		if err := rows.Scan(&user.Id, &user.Name, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.WithError(err).Error("Failed to scan a row")
		}

		res = append(res, &user)
	}

	return
}

func (r pgRepository) FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error) {
	return
}

func (r pgRepository) Delete(ctx context.Context, req DeleteRequest) (err error) {
	return
}
