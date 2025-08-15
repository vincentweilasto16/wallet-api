package repository

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/vincentweilasto16/wallet-api/internal/repository/postgres"
)

//go:generate mockgen -source=./postgres_repository.go -destination=./mock/postgres_repository_mock.go -package=mock
type IPostgresRepository interface {
	// User Repository
	GetUserByID(ctx context.Context, id uuid.UUID) (entity.User, error)
	UpdateUserBalance(ctx context.Context, arg entity.UpdateUserBalanceParams) error

	// Transaction Repository
	CreateTransaction(ctx context.Context, arg entity.CreateTransactionParams) error
	GetUserTransactions(ctx context.Context, userID uuid.UUID) ([]entity.Transaction, error)
}
