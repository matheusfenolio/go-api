package routers

import (
	"go-api/src/internal/customer"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	version1 := router.Group("/v1")
	customer.InitCustomerController(version1)
}
