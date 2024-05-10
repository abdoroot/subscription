package store_test

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/store"
	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCustomerStore(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Fatal("error connecting to the database:", err)
	}
	customerStore := store.NewCustomerStore(db)

	t.Run("CreateCustomer", func(t *testing.T) {
		customer := types.Customer{
			Type:        "individual",
			DisplayName: "John Doe",
			FirstName:   "John",
			LastName:    "Doe",
			CompanyName: "",
			Email:       "john.doe@example.com",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		}
		createdCustomer, err := customerStore.CreateCustomer(customer)
		assert.NoError(t, err)
		assert.NotZero(t, createdCustomer.ID)
		assert.Equal(t, customer.Type, createdCustomer.Type)
		assert.Equal(t, customer.DisplayName, createdCustomer.DisplayName)
		// Add more assertions as needed for other fields
	})

	t.Run("UpdateCustomer", func(t *testing.T) {
		updatedCustomer := types.Customer{
			ID:          1, // Assuming the customer with ID 1 exists
			Type:        "company",
			DisplayName: "Updated Company",
			FirstName:   "",
			LastName:    "",
			CompanyName: "Updated Company Name",
			Email:       "updated@example.com",
			UpdatedAt:   time.Now().UTC(),
		}
		err := customerStore.UpdateCustomer(updatedCustomer)
		assert.NoError(t, err)
	})

	t.Run("GetCustomerByID", func(t *testing.T) {
		id := 1 // Assuming the customer with ID 1 exists
		customer, err := customerStore.GetCustomerByID(id)
		assert.NoError(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, id, customer.ID)
	})

	t.Run("GetCustomersByType", func(t *testing.T) {
		customers, err := customerStore.GetCustomersByCompanyId(1)
		assert.NoError(t, err)
		assert.NotEmpty(t, customers)
		for _, customer := range customers {
			assert.Equal(t, "individual", customer.Type)
		}
	})

	t.Run("DeleteCustomerByID", func(t *testing.T) {
		id := 1 // Assuming the customer with ID 1 exists
		err := customerStore.DeleteCustomerByID(id)
		assert.NoError(t, err)
	})
}
