package main

import (
	"fmt"
	"go-api/src/customer"
	"go-api/src/database"
	"go-api/src/routers"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	version1 := router.Group("/v1")

	routers.Customers(version1)

	initDatabase()

	router.Run()
}

func initDatabase() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s",
		os.Getenv("HOST"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DATABSE"),
		os.Getenv("DB_PORT"),
		os.Getenv("SDSLMODE"),
		os.Getenv("TIMEZONE"))

	database.DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.DBConnection.AutoMigrate(&customer.Customer{})
}
