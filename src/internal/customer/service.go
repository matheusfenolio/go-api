package customer

func GetCustumers() ([]Customer, error) {
	customers, err := fetchCustomers()
	return customers, err
}

func GetCustomerById(id int) (Customer, error) {
	customer, err := fetchCustomer(id)

	return customer, err
}

func CreateCustomer(customerRequest CustomerResquest) (Customer, error) {
	newCustomer := convertRequestToEntity(customerRequest)
	customer, err := saveCustomer(newCustomer)
	return customer, err
}

func ChangeCustomer(id int, customerRequest CustomerResquest) error {
	customerToChange := convertRequestToEntity(customerRequest)
	return updateCustomer(id, customerToChange)
}

func DeleteCustomer(id int) error {
	return DeleteCustomer(id)
}

func convertRequestToEntity(customerRequest CustomerResquest) Customer {
	return Customer{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Email:     customerRequest.Email,
	}
}
