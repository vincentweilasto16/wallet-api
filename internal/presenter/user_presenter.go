package presenter

import (
	"github.com/guregu/null"
	entity "github.com/vincentweilasto16/wallet-api/internal/repository/postgres"
	"github.com/vincentweilasto16/wallet-api/internal/response"
)

func UserResponse(user *entity.User) *response.UserResponse {
	if user == nil {
		return nil
	}

	return &response.UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Balance:   user.Balance,
		CreatedAt: null.NewTime(user.CreatedAt.Time, user.CreatedAt.Valid),
		UpdatedAt: null.NewTime(user.UpdatedAt.Time, user.UpdatedAt.Valid),
		DeletedAt: null.NewTime(user.DeletedAt.Time, user.DeletedAt.Valid),
	}
}
