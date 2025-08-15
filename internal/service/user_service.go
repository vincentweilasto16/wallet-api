package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/vincentweilasto16/wallet-api/internal/errors"
	"github.com/vincentweilasto16/wallet-api/internal/repository"
	entity "github.com/vincentweilasto16/wallet-api/internal/repository/postgres"
)

//go:generate mockgen -destination=./mock/user_service_mock.go -package=mock github.com/vincentweilasto16/wallet-api/internal/service IUserService
type IUserService interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
}

type UserService struct {
	repo repository.IPostgresRepository
}

func NewUserService(repo repository.IPostgresRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {

	userID, err := uuid.Parse(id)
	if err != nil {
		log.Println("GetUserByID failed to parse user id with error: ", err.Error())
		return nil, errors.ErrBadRequest.New("invalid user id")
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("GetUserByID failed to get user by id with error: ", err.Error())
		return nil, errors.ErrNotFound.New("user not found")
	}

	return &user, nil
}
