package request

import (
	"github.com/gin-gonic/gin"
)

func SetBodyParams(c *gin.Context, target interface{}) error {
	return c.ShouldBindJSON(target)
}

func SetURIParams(c *gin.Context, target interface{}) error {
	return c.ShouldBindUri(target)
}

func SetQueryParams(c *gin.Context, target interface{}) error {
	return c.ShouldBindQuery(target)
}
