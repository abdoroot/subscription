package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
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
		Attachments:        []int{1, 2,3},
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
	store := NewInvoiceStore(db)
	param := types.Invoice{
		CompanyID:          1,
		OrderNumber:        "Example Order",
		InvoiceDate:        time.Now(),
		Subject:            "Example Subject",
		TermsAndConditions: "Example Terms",
		Attachments:        []int{1, 2},
	}
	invoice, err := store.CreateInvoice(param)
	return store, invoice, err
}
