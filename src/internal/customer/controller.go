package customer

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitCustomerController(router *gin.RouterGroup) {
	customersRoute := router.Group("/customer")

	customersRoute.GET("/", func(context *gin.Context) {
		response, err := GetCustumers()

		if err != nil {
			context.AbortWithStatus(404)
			return
		}

		context.JSON(200, convertCustomerListToCustomerResponseList(response))
	})

	customersRoute.GET("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		response, err := GetCustomerById(id)

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

		response, err := CreateCustomer(data)

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

		err = ChangeCustomer(id, data)

		if err != nil {
			context.AbortWithError(500, err)
			return
		}

		context.Status(200)
	})

	customersRoute.DELETE("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		err := DeleteCustomer(id)

		if err != nil {
			context.AbortWithError(500, err)
			return
		}

		context.Status(200)
	})
}

func convertCustomerListToCustomerResponseList(customers []Customer) []CustomerResponse {
	var customersResponse []CustomerResponse

	for _, item := range customers {
		customersResponse = append(customersResponse, item.ConvertToCustomerResponse())
	}

	return customersResponse
}
