package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentweilasto16/wallet-api/internal/constants"
	"github.com/vincentweilasto16/wallet-api/internal/controller"
)

func NewRouter(ctrl *controller.Controllers) *gin.Engine {
	r := gin.Default()

	// public api v1
	publicV1 := r.Group(constants.PublicAPIV1BasePath)
	RegisterRoutes(publicV1, ctrl)

	return r
}
