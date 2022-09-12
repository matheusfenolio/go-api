package customer

type CustomerRepository struct {
}

func (r *CustomerRepository) fetchCustomers() ([]Customer, error) {
	var customers []Customer
	result := DBConnection.Find(&customers)
	return customers, result.Error
}

func (r *CustomerRepository) fetchCustomer(id int) (Customer, error) {
	var customer Customer
	result := DBConnection.First(&customer, id)
	return customer, result.Error
}

func (r *CustomerRepository) saveCustomer(customer Customer) (Customer, error) {
	result := DBConnection.Create(&customer)
	return customer, result.Error
}

func (r *CustomerRepository) updateCustomer(id int, customer Customer) error {
	result := DBConnection.Model(&Customer{}).Where("id = ?", id).Updates(
		Customer{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
		})

	return result.Error
}

func (r *CustomerRepository) deleteCustomer(id int) error {
	result := DBConnection.Delete(&Customer{}, id)
	return result.Error
}
