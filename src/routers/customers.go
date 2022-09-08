package routers

import (
	customer "go-api/src/customer"

	"strconv"

	"github.com/gin-gonic/gin"
)

func Customers(router *gin.RouterGroup) {
	customersRoute := router.Group("/customer")

	customersRoute.GET("/", func(context *gin.Context) {
		response, err := customer.GetCustumers()

		if err != nil {
			context.AbortWithStatus(404)
			return
		}

		context.JSON(200, convertCustomerListToCustomerResponseList(response))
	})

	customersRoute.GET("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		response, err := customer.GetCustomerById(id)

		if err != nil {
			context.AbortWithStatus(404)
			return
		}

		context.JSON(200, response)
	})

	customersRoute.POST("/", func(context *gin.Context) {

		var data customer.CustomerResquest

		err := context.ShouldBind(&data)

		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": "Invalid body format"})
			return
		}

		response, err := customer.CreateCustomer(data)

		if err != nil {
			context.AbortWithStatus(500)
			return
		}

		context.JSON(200, response)
	})

	customersRoute.PUT("/:id", func(context *gin.Context) {

		var data customer.CustomerResquest

		err := context.ShouldBind(&data)

		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": "Invalid body format"})
			return
		}

		id, _ := strconv.Atoi(context.Param("id"))

		err = customer.ChangeCustomer(id, data)

		if err != nil {
			context.AbortWithError(500, err)
			return
		}

		context.Status(200)
	})

	customersRoute.DELETE("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		err := customer.DeleteCustomer(id)

		if err != nil {
			context.AbortWithError(500, err)
			return
		}

		context.Status(200)
	})
}

func convertCustomerListToCustomerResponseList(customers []customer.Customer) []customer.CustomerResponse {
	var customersResponse []customer.CustomerResponse

	for _, customer := range customers {
		customersResponse = append(customersResponse, customer.ConvertToCustomerResponse())
	}

	return customersResponse
}
