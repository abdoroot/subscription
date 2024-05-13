package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

var createdAddress types.Address

func TestCreateAddress(t *testing.T) {
	db, err := util.ConnectToPq()
	defer func() {
		db.Close()
	}()
	if err != nil {
		t.Error("error connect to db")
	}
	addressStore := NewAddressStore(db)

	address := types.Address{
		Type:       "shipping",
		CustomerID: 1,
		CountryID:  2,
		CityID:     3,
		Line1:      "123 Main St",
		Line2:      "Apt 101",
	}

	createdAddress, err = addressStore.CreateAddress(address)
	if err != nil {
		t.Errorf("error creating address: %v", err)
	}

	assert.Greater(t, createdAddress.ID, 0)
	assert.Equal(t, "shipping", createdAddress.Type)
	assert.Equal(t, 1, createdAddress.CustomerID)
	assert.Equal(t, 2, createdAddress.CountryID)
	assert.Equal(t, 3, createdAddress.CityID)
	assert.Equal(t, "123 Main St", createdAddress.Line1)
	assert.Equal(t, "Apt 101", createdAddress.Line2)
}

func TestUpdateAddress(t *testing.T) {
	if createdAddress.ID == 0 {
		t.Error("TestUpdateAddress depends on createdAddress from TestCreateAddress")
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	addressStore := NewAddressStore(db)

	updatedAddress := createdAddress
	updatedAddress.Type = "billing"
	updatedAddress.Line2 = "Suite 200"

	err = addressStore.UpdateAddress(updatedAddress)
	if err != nil {
		t.Errorf("error updating address: %v", err)
	}
}

func TestGetAddressByCustomerId(t *testing.T) {
	if createdAddress.ID == 0 {
		t.Error("TestGetAddressByCustomerId depends on createdAddress from TestCreateAddress")
		return
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connecting to the database")
		return
	}

	addressStore := NewAddressStore(db)
	addresses, err := addressStore.GetAddressByCustomerId(createdAddress.CustomerID)
	if err != nil {
		t.Errorf("error getting addresses by customer ID: %v", err)
		return
	}

	assert := assert.New(t)
	assert.NotEmpty(addresses, "no addresses found for customer ID")

	for _, address := range addresses {
		assert.Equal(createdAddress.CustomerID, address.CustomerID, "address does not belong to the correct customer ID")
	}
}

func TestDeleteAddressByID(t *testing.T) {
	if createdAddress.ID == 0 {
		t.Error("TestDeleteAddressByID depends on createdAddress from TestCreateAddress")
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	addressStore := NewAddressStore(db)

	err = addressStore.DeleteAddressByID(createdAddress.ID)
	assert.Nil(t, err, "DeleteAddressByID error")
}
