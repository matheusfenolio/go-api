package customer

type CustomerService struct {
}

func (s *CustomerService) getCustumers() ([]Customer, error) {
	customers, err := respository.fetchCustomers()
	return customers, err
}

func (s *CustomerService) getCustomerById(id int) (Customer, error) {
	customer, err := respository.fetchCustomer(id)

	return customer, err
}

func (s *CustomerService) createCustomer(customerToCreate Customer) (Customer, error) {
	customer, err := respository.saveCustomer(customerToCreate)
	return customer, err
}

func (s *CustomerService) updateCustomer(id int, customerToUpdate Customer) error {
	return respository.updateCustomer(id, customerToUpdate)
}

func (s *CustomerService) deleteCustomer(id int) error {
	return respository.deleteCustomer(id)
}
