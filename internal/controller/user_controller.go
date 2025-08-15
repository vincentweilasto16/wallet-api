package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentweilasto16/wallet-api/internal/presenter"
	"github.com/vincentweilasto16/wallet-api/internal/request"
	"github.com/vincentweilasto16/wallet-api/internal/response"
	"github.com/vincentweilasto16/wallet-api/internal/service"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(UserService service.IUserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	// @TODO: prepare the context

	var uriParams request.GetUserByIDRequest
	if err := request.SetURIParams(ctx, &uriParams); err != nil {
		response.Error(ctx, err)
		return
	}

	user, err := c.UserService.GetUserByID(ctx, uriParams.UserID)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	res := presenter.UserResponse(user)
	response.Success(ctx, res, "ok")
}
