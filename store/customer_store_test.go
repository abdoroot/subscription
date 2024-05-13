package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

var globalCustomer types.Customer

func TestCustomerStore(t *testing.T) {
	db, err := util.ConnectToPq()
	defer func() {
		db.Close()
	}()
	if err != nil {
		t.Fatal("error connecting to the database:", err)
	}
	addStore := NewAddressStore(db)
	customerStore := NewCustomerStore(db, addStore)

	t.Run("CreateCustomer", func(t *testing.T) {
		addr := types.Address{
			Type:      "shipping",
			CountryID: 2,
			CityID:    3,
			Line1:     "123 Main St",
			Line2:     "Apt 101",
		}
		customer := types.Customer{
			CompanyId:   1,
			Type:        "individual",
			DisplayName: "John Doe",
			FirstName:   "John",
			LastName:    "Doe",
			CompanyName: "",
			Email:       "john.doe@example.com",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Address: []types.Address{
				addr,
			},
		}
		createdCustomer, err := customerStore.CreateCustomer(customer)
		globalCustomer = createdCustomer
		assert.NoError(t, err)
		assert.NotZero(t, createdCustomer.ID)
		assert.Equal(t, customer.Type, createdCustomer.Type)
		assert.Equal(t, customer.DisplayName, createdCustomer.DisplayName)
	})

	t.Run("UpdateCustomer", func(t *testing.T) {
		if globalCustomer.ID == 0 {
			t.Error("test UpdateCustomer depend on globalCustomer")
		}
		updatedCustomer := types.Customer{
			ID:          globalCustomer.ID,
			Type:        "company",
			DisplayName: "Updated name",
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
		if globalCustomer.ID == 0 {
			t.Error("test GetCustomerByID depend on globalCustomer")
		}
		id := globalCustomer.ID
		customer, err := customerStore.GetCustomerByID(id)
		if err != nil {
			t.Fatal("error GetCustomerByID :", err)
		}
		assert.NotNil(t, customer)
		assert.Equal(t, id, customer.ID)
	})

	t.Run("GetCustomersByCompanyId", func(t *testing.T) {
		if globalCustomer.ID == 0 {
			t.Error("test GetCustomersByCompanyId depend on globalCustomer")
		}
		customers, err := customerStore.GetCustomersByCompanyId(globalCustomer.CompanyId)
		if err != nil {
			t.Fatal("error GetCustomersByCompanyId :", err)
		}
		assert.NotEmpty(t, customers)
		assert.Greater(t, len(customers), 0)
	})

	t.Run("DeleteCustomerByID", func(t *testing.T) {
		if globalCustomer.ID == 0 {
			t.Error("test DeleteCustomerByID depend on globalCustomer")
		}
		id := globalCustomer.ID
		err := customerStore.DeleteCustomerByID(id)
		if err != nil {
			t.Error("error DeleteCustomerByID:", err)
		}
		assert.NoError(t, err)
	})
}

func TestGetAll(t *testing.T) {
	cs, c, err := createCustomer()
	if err != nil {
		t.Error("err createCustomer")
	}
	defer func() {
		if c.ID > 0 {
			cs.DeleteCustomerByID(c.ID)
		}
		cs.db.Close()
	}()
	//test
	customers, err := cs.GetAll()
	if err != nil {
		t.Error("err GetAll customers")
	}
	assert.Greater(t, len(customers), 0)
}

// helper func
func createCustomer() (*customerStore, types.Customer, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.Customer{}, nil
	}
	addStore := NewAddressStore(db)
	customerStore := NewCustomerStore(db, addStore)
	addr := types.Address{
		Type:      "shipping",
		CountryID: 2,
		CityID:    3,
		Line1:     "123 Main St",
		Line2:     "Apt 101",
	}
	customer := types.Customer{
		CompanyId:   1,
		Type:        "individual",
		DisplayName: "John Doe",
		FirstName:   "John",
		LastName:    "Doe",
		CompanyName: "",
		Email:       "john.doe@example.com",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Address: []types.Address{
			addr,
		},
	}
	createdCustomer, err := customerStore.CreateCustomer(customer)
	if err != nil {
		return nil, types.Customer{}, nil
	}

	return customerStore, createdCustomer, nil
}