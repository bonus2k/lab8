package services

import (
	"context"
	"github.com/bonus2k/lab8/internal/models"
	"github.com/bonus2k/lab8/internal/repositories"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.UserReq) (*models.UserRes, error)
	GetUser(ctx context.Context, id uuid.UUID) (*models.UserRes, error)
	GetAllUsers(ctx context.Context) (*[]models.UserRes, error)
}

var (
	service UserService
)

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func (u UserServiceImpl) CreateUser(ctx context.Context, user models.UserReq) (*models.UserRes, error) {
	entity := models.User{}
	entity.ToEntity(user)
	err := u.repo.CreateUser(ctx, &entity)
	if err != nil {
		return nil, err
	}
	res := models.UserRes{}
	res.ToDto(entity)
	return &res, nil
}

func (u UserServiceImpl) GetUser(ctx context.Context, id uuid.UUID) (*models.UserRes, error) {
	entity, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	res := models.UserRes{}
	res.ToDto(*entity)
	return &res, nil
}

func (u UserServiceImpl) GetAllUsers(ctx context.Context) (*[]models.UserRes, error) {
	entities, err := u.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]models.UserRes, len(entities))
	for i, entity := range entities {
		res[i].ToDto(entity)
	}
	return &res, nil
}

func Init(repo *repositories.UserRepository) UserService {
	service = UserServiceImpl{repo: *repo}
	return service
}
