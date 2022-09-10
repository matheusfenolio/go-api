package customer

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitCustomerController(router *gin.RouterGroup) {
	customersRoute := router.Group("/customer")

	customersRoute.GET("/", func(context *gin.Context) {
		response, err := service.getCustumers()

		if err != nil {
			context.AbortWithStatus(404)
			return
		}

		context.JSON(200, mapper.convertCustomerListToCustomerResponseList(response))
	})

	customersRoute.GET("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		response, err := service.getCustomerById(id)

		if err != nil {
			context.AbortWithStatus(404)
			return
		}

		context.JSON(200, response)
	})

	customersRoute.POST("/", func(context *gin.Context) {

		var data CustomerResquest

		err := context.ShouldBind(&data)

		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": "Invalid body format"})
			return
		}

		response, err := service.createCustomer(mapper.convertRequestToEntity(data))

		if err != nil {
			context.AbortWithStatus(500)
			return
		}

		context.JSON(200, response)
	})

	customersRoute.PUT("/:id", func(context *gin.Context) {

		var data CustomerResquest

		err := context.ShouldBind(&data)

		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": "Invalid body format"})
			return
		}

		id, _ := strconv.Atoi(context.Param("id"))

		err = service.updateCustomer(id, mapper.convertRequestToEntity(data))

		if err != nil {
			context.AbortWithError(500, err)
			return
		}

		context.Status(200)
	})

	customersRoute.DELETE("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		err := service.deleteCustomer(id)

		if err != nil {
			context.AbortWithError(500, err)
			return
		}

		context.Status(200)
	})
}
