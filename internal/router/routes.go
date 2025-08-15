package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentweilasto16/wallet-api/internal/constants"
	"github.com/vincentweilasto16/wallet-api/internal/controller"
)

func RegisterRoutes(rg *gin.RouterGroup, ctrl *controller.Controllers) {

	// users route
	users := rg.Group(constants.UserBasePath)
	{
		users.GET("/:id", ctrl.UserController.GetUserByID)
	}

	// transactions route
	tx := rg.Group(constants.TransactionBasePath)
	{
		tx.POST("/withdraw", ctrl.TransactionController.Withdraw)
	}
}
