package main

import (
	"go-api/src/internal/database"
	"go-api/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.InitDatabase()

	routers.InitRoutes(router)

	router.Run()
}
