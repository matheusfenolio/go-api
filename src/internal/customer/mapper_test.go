package customer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToCustomerResponse(t *testing.T) {

	entity := Customer{
		FirstName: "Jhon",
		LastName:  "Contoso",
		Email:     "jhon.c@email.com",
	}

	mapper := CustomerMapper{}

	result := mapper.convertToCustomerResponse(entity)

	require.Equal(t, "Jhon", result.FirstName)
	require.Equal(t, "Contoso", result.LastName)
	require.Equal(t, "jhon.c@email.com", result.Email)
}

func TestConvertCustomerListToCustomerResponseList(t *testing.T) {

	entity := Customer{
		FirstName: "Jhon",
		LastName:  "Contoso",
		Email:     "jhon.c@email.com",
	}

	entities := []Customer{entity}

	mapper := CustomerMapper{}

	result := mapper.convertCustomerListToCustomerResponseList(entities)

	require.NotEmpty(t, result)
	require.Equal(t, "Jhon", result[0].FirstName)
	require.Equal(t, "Contoso", result[0].LastName)
	require.Equal(t, "jhon.c@email.com", result[0].Email)
}

func TestConvertRequestToEntity(t *testing.T) {

	request := CustomerResquest{
		FirstName: "Jhon",
		LastName:  "Contoso",
		Email:     "jhon.c@email.com",
	}

	mapper := CustomerMapper{}

	result := mapper.convertRequestToEntity(request)

	require.NotEmpty(t, result)
	require.Equal(t, "Jhon", result.FirstName)
	require.Equal(t, "Contoso", result.LastName)
	require.Equal(t, "jhon.c@email.com", result.Email)
}
