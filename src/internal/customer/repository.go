package customer

import (
	"go-api/src/internal/database"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

func (customer Customer) ConvertToCustomerResponse() CustomerResponse {
	return CustomerResponse{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}
}

func fetchCustomers() ([]Customer, error) {
	var customers []Customer
	result := database.DBConnection.Find(&customers)
	return customers, result.Error
}

func fetchCustomer(id int) (Customer, error) {
	var customer Customer
	result := database.DBConnection.First(&customer, id)
	return customer, result.Error
}

func saveCustomer(customer Customer) (Customer, error) {
	result := database.DBConnection.Create(&customer)
	return customer, result.Error
}

func updateCustomer(id int, customer Customer) error {
	result := database.DBConnection.Model(&Customer{}).Where("id = ?", id).Updates(
		Customer{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
		})

	return result.Error
}

func deleteCustomer(id int) error {
	result := database.DBConnection.Delete(&Customer{}, id)
	return result.Error
}
