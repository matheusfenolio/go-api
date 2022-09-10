package database

import (
	"fmt"
	"go-api/src/internal/customer"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConnection *gorm.DB
)

func InitDatabase() {
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

	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	injectDependency()
}

func injectDependency() {
	customer.InitCustomerRepository(DBConnection)
}
