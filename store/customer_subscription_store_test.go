package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomerSubscription(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}
	defer func() {
		if subscription.ID > 0 {
			store.DeleteCustomerSubscriptionByID(subscription.ID)
		}
		store.db.Close()
	}()
	assert.NoError(t, err)
	assert.Greater(t, subscription.ID, 0)
	assert.NotZero(t, subscription.CreatedAt)
}

func TestUpdateCustomerSubscription(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}
	defer func() {
		if subscription.ID > 0 {
			store.DeleteCustomerSubscriptionByID(subscription.ID)
		}
		store.db.Close()
	}()

	param := types.CustomerSubscription{
		CompanyID:        1,
		CustomerID:       2,
		PlanID:           1,
		StartDate:        time.Now(),
		NeverExpire:      false,
		ExpireAfterCycle: 12,
		Qty:              1,
		Price:            99.99,
		Amount:           99.99,
		Reference:        "#123",
		OfflinePayment:   false,
	}

	err = store.UpdateCustomerSubscriptionByID(param, subscription.ID)
	assert.NoError(t, err)

	updatedSubscription, err := store.GetCustomerSubscriptionByID(subscription.ID)
	assert.NoError(t, err)
	assert.Equal(t, param.CustomerID, updatedSubscription.CustomerID)
	assert.Equal(t, param.PlanID, updatedSubscription.PlanID)
	assert.Equal(t, param.Price, updatedSubscription.Price)
}

func TestGetCustomerSubscriptionByID(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	defer func() {
		if subscription.ID > 0 {
			store.DeleteCustomerSubscriptionByID(subscription.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}

	got, err := store.GetCustomerSubscriptionByID(subscription.ID)
	assert.NoError(t, err)
	assert.Equal(t, subscription.ID, got.ID)
}

func TestGetCustomerSubscriptionsByCustomerID(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	defer func() {
		if subscription.ID > 0 {
			store.DeleteCustomerSubscriptionByID(subscription.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}

	subscriptions, err := store.GetCustomerSubscriptionsByCustomerID(subscription.CustomerID)
	assert.NoError(t, err)
	assert.Greater(t, len(subscriptions), 0)
}

func TestGetCustomerSubscriptionsByCompanyID(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	defer func() {
		if subscription.ID > 0 {
			store.DeleteCustomerSubscriptionByID(subscription.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}

	subscriptions, err := store.GetCustomerSubscriptionsByCompanyID(subscription.CompanyID)
	assert.NoError(t, err)
	assert.Greater(t, len(subscriptions), 0)
}

func TestGetAllCustomerSubscriptions(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	defer func() {
		if subscription.ID > 0 {
			store.DeleteCustomerSubscriptionByID(subscription.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}

	subscriptions, err := store.GetAllCustomerSubscriptions()
	assert.NoError(t, err)
	assert.Greater(t, len(subscriptions), 0)
}

func TestDeleteCustomerSubscriptionByID(t *testing.T) {
	store, subscription, err := CreateCustomerSubscription()
	if err != nil {
		t.Error("error creating customer subscription:", err)
	}
	defer store.db.Close()

	err = store.DeleteCustomerSubscriptionByID(subscription.ID)
	assert.NoError(t, err)

	_, err = store.GetCustomerSubscriptionByID(subscription.ID)
	assert.Error(t, err)
}

func CreateCustomerSubscription() (*CustomerSubscriptionStore, types.CustomerSubscription, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.CustomerSubscription{}, err
	}
	store := NewCustomerSubscriptionStore(db)
	subscription := types.CustomerSubscription{
		CompanyID:        1,
		CustomerID:       1,
		PlanID:           1,
		StartDate:        time.Now(),
		NeverExpire:      false,
		ExpireAfterCycle: 12,
		Qty:              1,
		Price:            99.99,
		Amount:           99.99,
		Reference:        "#123",
		OfflinePayment:   false,
	}
	subscription, err = store.CreateCustomerSubscription(subscription)
	return store, subscription, err
}
