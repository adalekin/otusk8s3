package userrepo

import (
	"context"

	"github.com/adalekin/otusk8s3/internal/domain"
)

type IRepository interface {
	Create(ctx context.Context, req CreateRequest) (res *domain.User, err error)
	FindAll(ctx context.Context) (res []*domain.User, err error)
	FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error)
	Update(ctx context.Context, req UpdateRequest) (res *domain.User, err error)
}

type CreateRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type FindByIdRequest struct {
	Id int64 `json:"id"`
}

type UpdateRequest struct {
	Id       int64   `json:"id"`
	Name     *string `json:"name,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}
