package userrepo

import (
	"context"
	"database/sql"
	"time"

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
	now := time.Now().UTC()

	query := "INSERT INTO users (name, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id;"
	res = &domain.User{Name: req.Name, IsActive: true, CreatedAt: now, UpdatedAt: now}

	err = r.db.QueryRow(query, res.Name, res.IsActive, res.CreatedAt, res.UpdatedAt).Scan(&res.Id)

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

	res = make([]*domain.User, 0)

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
	res = &domain.User{}

	query := "SELECT * FROM users WHERE id=$1;"
	err = r.db.QueryRow(query, req.Id).Scan(&res.Id, &res.Name, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		log.WithError(err).Error("Failed to scan a row")
		return nil, err
	}

	return
}

func (r pgRepository) Update(ctx context.Context, req UpdateRequest) (res *domain.User, err error) {
	now := time.Now().UTC()

	res, err = r.FindById(ctx, FindByIdRequest{Id: req.Id})

	if res == nil || err != nil {
		return res, err
	}

	if req.Name != nil {
		res.Name = *req.Name
	}

	if req.IsActive != nil {
		res.IsActive = *req.IsActive
	}

	res.UpdatedAt = now

	if err != nil {
		return nil, err
	}

	query := "UPDATE users SET name=COALESCE($2, name), is_active=COALESCE($3, is_active), updated_at=$4 WHERE id=$1;"

	_, err = r.db.Exec(query, res.Id, res.Name, res.IsActive, now)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		log.WithError(err).Error("Failed to scan a row")
		return nil, err
	}

	return
}
