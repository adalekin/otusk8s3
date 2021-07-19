package userrepo

import (
	"context"

	"github.com/adalekin/otusk8s3/internal/domain"
)

type IRepository interface {
	Create(ctx context.Context, req CreateRequest) (res *domain.User, err error)
	FindAll(ctx context.Context) (res []*domain.User, err error)
	FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error)
	Delete(ctx context.Context, req DeleteRequest) (err error)
}

type CreateRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:is_active`
}

type FindByIdRequest struct {
	Id int64 `json:"id"`
}

type DeleteRequest struct {
	Id int64 `json:"id"`
}
