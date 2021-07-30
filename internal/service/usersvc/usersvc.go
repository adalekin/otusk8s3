package usersvc

import (
	"context"

	"github.com/adalekin/otusk8s3/internal/domain"
	"github.com/adalekin/otusk8s3/internal/repository/reporegistry"
	"github.com/adalekin/otusk8s3/internal/repository/userrepo"
	"github.com/adalekin/otusk8s3/internal/util"
)

type IService interface {
	Create(ctx context.Context, req CreateRequest) (user *domain.User, err error)
	FindAll(ctx context.Context) (res []*domain.User, err error)
	FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error)
	Update(ctx context.Context, req UpdateRequest) (res *domain.User, err error)
	Delete(ctx context.Context, req DeleteRequest) (res *domain.User, err error)
}

func New(repoRegistry reporegistry.IRepoRegistry) IService {
	return &Service{
		repoRegistry: repoRegistry,
	}
}

type CreateRequest struct {
	Name string `json:"name"`
}

func (r CreateRequest) toCreateRepoReq() userrepo.CreateRequest {
	return userrepo.CreateRequest{
		Name: r.Name,
	}
}

type FindByIdRequest struct {
	Id int64 `json:"id"`
}

func (r FindByIdRequest) toFindByIdRepoReq() userrepo.FindByIdRequest {
	return userrepo.FindByIdRequest{
		Id: r.Id,
	}
}

type UpdateRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (r UpdateRequest) toUpdateRepoReq() userrepo.UpdateRequest {
	return userrepo.UpdateRequest{
		Id:   r.Id,
		Name: util.String(r.Name),
	}
}

type DeleteRequest struct {
	Id int64 `json:"id"`
}

func (r DeleteRequest) toUpdateRepoReq() userrepo.UpdateRequest {
	return userrepo.UpdateRequest{
		Id:       r.Id,
		IsActive: util.Bool(false),
	}
}

type Service struct {
	repoRegistry reporegistry.IRepoRegistry
}

func (s Service) Create(ctx context.Context, req CreateRequest) (res *domain.User, err error) {
	var createUser = func(repoRegistry reporegistry.IRepoRegistry) (interface{}, error) {
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
	var listUsers = func(repoRegistry reporegistry.IRepoRegistry) (interface{}, error) {
		userRepo := repoRegistry.GetUserRepository()
		users, err := userRepo.FindAll(ctx)

		if err != nil {
			return nil, err
		}

		return users, nil
	}

	var out interface{}
	out, err = s.repoRegistry.DoInTransaction(listUsers)

	if err != nil {
		return nil, err
	}

	res = out.([]*domain.User)
	return
}

func (s Service) FindById(ctx context.Context, req FindByIdRequest) (res *domain.User, err error) {
	var findUserById = func(repoRegistry reporegistry.IRepoRegistry) (interface{}, error) {
		userRepo := repoRegistry.GetUserRepository()
		user, err := userRepo.FindById(ctx, req.toFindByIdRepoReq())

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	var out interface{}
	out, err = s.repoRegistry.DoInTransaction(findUserById)

	if err != nil {
		return nil, err
	}

	res = out.(*domain.User)
	return
}

func (s Service) Update(ctx context.Context, req UpdateRequest) (res *domain.User, err error) {
	var updateUser = func(repoRegistry reporegistry.IRepoRegistry) (interface{}, error) {
		userRepo := repoRegistry.GetUserRepository()

		var user *domain.User
		user, err = userRepo.Update(ctx, req.toUpdateRepoReq())

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	var out interface{}
	out, err = s.repoRegistry.DoInTransaction(updateUser)

	if err != nil {
		return nil, err
	}

	res = out.(*domain.User)
	return
}

func (s Service) Delete(ctx context.Context, req DeleteRequest) (res *domain.User, err error) {
	var updateUser = func(repoRegistry reporegistry.IRepoRegistry) (interface{}, error) {
		userRepo := repoRegistry.GetUserRepository()

		var user *domain.User
		user, err = userRepo.Update(ctx, req.toUpdateRepoReq())

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	var out interface{}
	out, err = s.repoRegistry.DoInTransaction(updateUser)

	if err != nil {
		return nil, err
	}

	res = out.(*domain.User)
	return
}
