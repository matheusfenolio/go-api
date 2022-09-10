package customer

import "gorm.io/gorm"

var (
	service      CustomerService    = CustomerService{}
	respository  CustomerRepository = CustomerRepository{}
	mapper       CustomerMapper     = CustomerMapper{}
	DBConnection *gorm.DB
)

func InitCustomerRepository(db *gorm.DB) {
	DBConnection = db
	DBConnection.AutoMigrate(&Customer{})
}
