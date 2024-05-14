package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	store, invoice, err := CreateInvoice()
	if err != nil {
		t.Fatal("error creating invoice:", err)
	}
	defer func() {
		if invoice.ID > 0 {
			store.DeleteInvoiceByID(invoice.ID)
		}
		store.db.Close()
	}()
	assert.NoError(t, err)
	assert.Greater(t, invoice.ID, 0)
	assert.NotZero(t, invoice.CreatedAt)
}

func TestGetInvoiceByID(t *testing.T) {
	store, invoice, err := CreateInvoice()
	defer func() {
		if invoice.ID > 0 {
			store.DeleteInvoiceByID(invoice.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Fatal("error creating invoice:", err)
	}

	got, err := store.GetInvoiceByID(invoice.ID)
	assert.NoError(t, err)
	assert.Equal(t, invoice.ID, got.ID)
}

func TestUpdateInvoiceByID(t *testing.T) {
	store, invoice, err := CreateInvoice()
	defer func() {
		if invoice.ID > 0 {
			store.DeleteInvoiceByID(invoice.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Fatal("error creating invoice:", err)
	}

	param := types.Invoice{
		CompanyID:          1,
		OrderNumber:        "updated order number",
		InvoiceDate:        time.Now(),
		Subject:            "updated subject",
		TermsAndConditions: "updated terms",
		Attachments:        pq.Int32Array{1, 2, 3},
	}

	err = store.UpdateInvoiceByID(param, invoice.ID)
	assert.NoError(t, err)

	updatedInvoice, err := store.GetInvoiceByID(invoice.ID)
	assert.NoError(t, err)
	assert.Equal(t, param.OrderNumber, updatedInvoice.OrderNumber)
	assert.Equal(t, param.Subject, updatedInvoice.Subject)
}

func TestDeleteInvoiceByID(t *testing.T) {
	store, invoice, err := CreateInvoice()
	if err != nil {
		t.Fatal("error creating invoice:", err)
	}
	defer store.db.Close()

	err = store.DeleteInvoiceByID(invoice.ID)
	assert.NoError(t, err)

	_, err = store.GetInvoiceByID(invoice.ID)
	assert.Error(t, err)
}

func TestGetInvoicesByCompanyID(t *testing.T) {
	store, invoice, err := CreateInvoice()
	defer func() {
		if invoice.ID > 0 {
			store.DeleteInvoiceByID(invoice.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Fatal("error creating invoice:", err)
	}

	invoices, err := store.GetInvoicesByCompanyID(invoice.CompanyID)
	assert.NoError(t, err)
	assert.Greater(t, len(invoices), 0)
}

func CreateInvoice() (*InvoiceStore, types.Invoice, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.Invoice{}, err
	}
	itemStore := NewInvoiceItemStore(db)
	store := NewInvoiceStore(db, itemStore)
	param := types.Invoice{
		CompanyID:          1,
		CustomerID:         2,
		OrderNumber:        "Example Order",
		InvoiceDate:        time.Now(),
		DueDate:            time.Now().Add((time.Hour * 24) * 3), //three days
		Subject:            "Example Subject",
		TermsAndConditions: "Example Terms",
		Attachments:        pq.Int32Array{1, 2},
		Items: []types.InvoiceItem{
			{
				CompanyID: 1,
				Type:      "product",
				Qty:       2,
				Rate:      10.00,
				ItemID:    123,
				Discount:  0.00,
				Amount:    10.00,
			},
		},
	}
	invoice, err := store.CreateInvoice(param)
	return store, invoice, err
}
