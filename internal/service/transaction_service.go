package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/vincentweilasto16/wallet-api/internal/constants"
	"github.com/vincentweilasto16/wallet-api/internal/errors"
	"github.com/vincentweilasto16/wallet-api/internal/repository"
	entity "github.com/vincentweilasto16/wallet-api/internal/repository/postgres"
	"github.com/vincentweilasto16/wallet-api/internal/request"
)

//go:generate mockgen -destination=./mock/transaction_service_mock.go -package=mock github.com/vincentweilasto16/wallet-api/internal/service ITransactionService
type ITransactionService interface {
	Withdraw(ctx context.Context, arg *request.WithdrawRequest) error
}

type TransactionService struct {
	repo repository.IPostgresRepository
}

func NewTransactionService(repo repository.IPostgresRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) Withdraw(ctx context.Context, params *request.WithdrawRequest) error {

	userID, err := uuid.Parse(params.UserID)
	if err != nil {
		log.Println("Withdraw failed to parse user id with error: ", err.Error())
		return errors.ErrBadRequest.New("invalid user id")
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("Withdraw failed to get user by id with error: ", err.Error())
		return errors.ErrNotFound.New("user not found")
	}

	if user.Balance < params.Amount {
		log.Println("Withdraw failed due to insuficient balance")
		return errors.ErrUnprocessableEntity.New("insuficient balance")
	}

	if err = s.repo.CreateTransaction(ctx, entity.CreateTransactionParams{
		UserID: user.ID,
		Amount: params.Amount,
		Type:   constants.TransactionTypeWithdraw,
		Status: constants.TransactionStatusCompleted,
		Description: sql.NullString{
			String: "withdraw a fund from wallet",
			Valid:  true,
		},
	}); err != nil {
		log.Println("Withdraw failed to create transaction with error: ", err.Error())
		return errors.ErrInternalServer.New("failed to create transaction record")
	}

	newBalance := user.Balance - params.Amount
	if err := s.repo.UpdateUserBalance(ctx, entity.UpdateUserBalanceParams{
		Balance: newBalance,
		ID:      user.ID,
	}); err != nil {
		log.Println("Withdraw failed to update user balance with error: ", err.Error())
		return errors.ErrInternalServer.New("failed to update user balance")
	}

	return nil
}
