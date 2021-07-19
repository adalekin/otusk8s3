package usersvc

import (
	"context"

	"github.com/adalekin/otusk8s3/internal/domain"
	"github.com/adalekin/otusk8s3/internal/repository/reporegistry"
	"github.com/adalekin/otusk8s3/internal/repository/userrepo"
)

type IService interface {
	Create(ctx context.Context, req CreateRequest) (user *domain.User, err error)
	FindAll(ctx context.Context) (res []*domain.User, err error)
	FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error)
	Delete(ctx context.Context, req DeleteRequest) (err error)
}

func New(repoRegistry reporegistry.RepoRegistry) IService {
	return &Service{
		repoRegistry: repoRegistry,
	}
}

type CreateRequest struct {
	Name string `json:"name"`
}

type FindByIdRequest struct {
	Id int64 `json:"id"`
}

type DeleteRequest struct {
	Id int64 `json:"id"`
}

func (r CreateRequest) toCreateRepoReq() userrepo.CreateRequest {
	return userrepo.CreateRequest{
		Name: r.Name,
	}
}

type Service struct {
	repoRegistry reporegistry.RepoRegistry
}

func (s Service) Create(ctx context.Context, req CreateRequest) (res *domain.User, err error) {
	var createUser = func(repoRegistry reporegistry.RepoRegistry) (interface{}, error) {
		userRepo := repoRegistry.GetUserRepository()

		var user *domain.User
		user, err = userRepo.Create(ctx, req.toCreateRepoReq())

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	var out interface{}
	out, err = s.repoRegistry.DoInTransaction(createUser)

	if err != nil {
		return nil, err
	}

	res = out.(*domain.User)
	return
}

func (s Service) FindAll(ctx context.Context) (res []*domain.User, err error) {
}
