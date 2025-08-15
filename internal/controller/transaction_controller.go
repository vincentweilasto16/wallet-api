package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentweilasto16/wallet-api/internal/request"
	"github.com/vincentweilasto16/wallet-api/internal/response"
	"github.com/vincentweilasto16/wallet-api/internal/service"
)

type TransactionController struct {
	TransactionService service.ITransactionService
}

func NewTransactionController(txSvc service.ITransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: txSvc,
	}
}

func (c *TransactionController) Withdraw(ctx *gin.Context) {
	// @TODO: prepare the context

	var req request.WithdrawRequest
	if err := request.SetBodyParams(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := c.TransactionService.Withdraw(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, gin.H{"message": "Withdrawal successful"}, "OK")
}
