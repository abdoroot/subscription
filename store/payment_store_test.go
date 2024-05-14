package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreatePayment(t *testing.T) {
	store, payment, err := CreatePayment()
	if err != nil {
		t.Error("err CreatePayment ", err)
	}
	defer func() {
		if payment.ID > 0 {
			store.DeletePaymentByID(payment.ID)
		}
	}()
	assert.Nil(t, err)
	assert.Greater(t, payment.ID, 0)
	assert.Greater(t, payment.CreatedAt, time.Time{})
}

func TestUpdatePayment(t *testing.T) {
	store, payment, err := CreatePayment()
	if err != nil {
		t.Error("err CreatePayment ", err)
	}
	defer func() {
		if payment.ID > 0 {
			store.DeletePaymentByID(payment.ID)
		}
	}()

	param := types.Payment{
		CompanyID:       1,
		CustomerID:      2,
		InvoiceID:       2,
		AmountReceived:  9.99,
		PaymentDate:     time.Now(),
		PaymentMethodID: 1,
		Reference:       "#225",
		Attachments:     pq.Int32Array{1, 2},
	}

	err = store.UpdatePaymentByID(param, payment.ID)
	assert.Nil(t, err)

	updatedPayment, err := store.GetPaymentByID(payment.ID)
	assert.Nil(t, err)
	assert.Equal(t, param.CustomerID, updatedPayment.CustomerID)
	assert.Equal(t, param.InvoiceID, updatedPayment.InvoiceID)
	assert.Equal(t, param.AmountReceived, updatedPayment.AmountReceived)
	assert.Equal(t, param.Reference, updatedPayment.Reference)
}

func TestGetPayment(t *testing.T) {
	store, payment, err := CreatePayment()
	if err != nil {
		t.Error("err CreatePayment ", err)
	}
	defer func() {
		if payment.ID > 0 {
			store.DeletePaymentByID(payment.ID)
		}
	}()

	t.Run("GetPaymentsByInvoiceID", func(t *testing.T) {
		ps, err := store.GetPaymentsByInvoiceID(payment.InvoiceID)
		assert.Nil(t, err)
		assert.Greater(t, len(ps), 0)
		assert.Greater(t, ps[0].ID, 0)
	})

	t.Run("GetPaymentsByCompanyID", func(t *testing.T) {
		ps, err := store.GetPaymentsByCompanyID(payment.CompanyID)
		assert.Nil(t, err)
		assert.Greater(t, len(ps), 0)
		assert.Greater(t, ps[0].ID, 0)
	})

	t.Run("GetPaymentsByCustomerID", func(t *testing.T) {
		ps, err := store.GetPaymentsByCustomerID(payment.CustomerID)
		assert.Nil(t, err)
		assert.Greater(t, len(ps), 0)
		assert.Greater(t, ps[0].ID, 0)
	})

	t.Run("GetAllPayments", func(t *testing.T) {
		ps, err := store.GetPaymentsByCustomerID(payment.CompanyID)
		assert.Nil(t, err)
		assert.Greater(t, len(ps), 0)
		assert.Greater(t, ps[0].ID, 0)
	})
}

func TestDeletePaymentById(t *testing.T) {
	store, payment, err := CreatePayment()
	if err != nil {
		t.Error("err CreatePayment ", err)
	}
	assert.Nil(t, err)
	err = store.DeletePaymentByID(payment.ID)
	assert.Nil(t, err)
}

func CreatePayment() (*PaymentStore, types.Payment, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.Payment{}, err
	}
	store := NewPaymentStore(db)
	payment := types.Payment{
		CompanyID:       1,
		CustomerID:      1,
		InvoiceID:       1,
		AmountReceived:  100.0,
		PaymentDate:     time.Now(),
		PaymentMethodID: 1,
		Reference:       "#225",
		Attachments:     pq.Int32Array{1, 2},
	}
	p, err := store.CreatePayment(payment)
	return store, p, err
}
