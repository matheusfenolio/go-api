package customer

type CustomerMapper struct {
}

func (m *CustomerMapper) convertCustomerListToCustomerResponseList(customers []Customer) []CustomerResponse {
	var customersResponse []CustomerResponse

	for _, item := range customers {
		customersResponse = append(customersResponse, m.convertToCustomerResponse(item))
	}

	return customersResponse
}

func (m *CustomerMapper) convertToCustomerResponse(customer Customer) CustomerResponse {
	return CustomerResponse{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}
}

func (m *CustomerMapper) convertRequestToEntity(customerRequest CustomerResquest) Customer {
	return Customer{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Email:     customerRequest.Email,
	}
}
