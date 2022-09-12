package database

import (
	"go-api/src/internal/customer"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConnection *gorm.DB
)

func InitDatabase() {
	var err error

	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s",
	// 	os.Getenv("HOST"),
	// 	os.Getenv("USERNAME"),
	// 	os.Getenv("PASSWORD"),
	// 	os.Getenv("DATABSE"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("SDSLMODE"),
	// 	os.Getenv("TIMEZONE"))

	dsn := "host=192.168.1.16 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"

	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	injectDependency()
}

func injectDependency() {
	customer.InitCustomerRepository(DBConnection)
}
