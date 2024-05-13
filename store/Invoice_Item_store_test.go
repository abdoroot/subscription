package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoiceItem(t *testing.T) {
	store, invoiceItem, err := CreateInvoiceItem()
	defer func() {
		if invoiceItem.ID > 0 {
			store.DeleteInvoiceItemByID(invoiceItem.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Fatal("error creating invoice item:", err)
	}
	assert.NoError(t, err)
	assert.Greater(t, invoiceItem.ID, 0)
	assert.NotZero(t, invoiceItem.CreatedAt)
}

func TestGetInvoiceItemByID(t *testing.T) {
	store, invoiceItem, err := CreateInvoiceItem()
	defer func() {
		if invoiceItem.ID > 0 {
			store.DeleteInvoiceItemByID(invoiceItem.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Fatal("error creating invoice item:", err)
	}

	got, err := store.GetInvoiceItemByID(invoiceItem.ID)
	assert.NoError(t, err)
	assert.Equal(t, invoiceItem.ID, got.ID)
	assert.Equal(t, invoiceItem.CompanyID, got.CompanyID)
}

func TestUpdateInvoiceItemByID(t *testing.T) {
	store, invoiceItem, err := CreateInvoiceItem()
	defer func() {
		if invoiceItem.ID > 0 {
			store.DeleteInvoiceItemByID(invoiceItem.ID)
		}
		store.db.Close()
	}()
	if err != nil {
		t.Fatal("error creating invoice item:", err)
	}

	param := types.InvoiceItem{
		CompanyID: 2,
		InvoiceID: 3,
		Type:      "product",
		Qty:       2,
		Rate:      11.00,
		ItemID:    124,
		Discount:  0.00,
		Amount:    12.00,
	}

	err = store.UpdateInvoiceItemByID(param, invoiceItem.ID)
	assert.NoError(t, err)

	updatedItem, err := store.GetInvoiceItemByID(invoiceItem.ID)
	assert.NoError(t, err)
	assert.Equal(t, param.Qty, updatedItem.Qty)
	assert.Equal(t, param.Rate, updatedItem.Rate)
}

func TestDeleteInvoiceItemByID(t *testing.T) {
	store, invoiceItem, err := CreateInvoiceItem()
	if err != nil {
		t.Fatal("error creating invoice item:", err)
	}
	defer store.db.Close()

	err = store.DeleteInvoiceItemByID(invoiceItem.ID)
	assert.NoError(t, err)

	_, err = store.GetInvoiceItemByID(invoiceItem.ID)
	assert.Error(t, err)
}

func CreateInvoiceItem() (*invoiceItemStore, types.InvoiceItem, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.InvoiceItem{}, err
	}
	store := NewInvoiceItemStore(db)
	param := types.InvoiceItem{
		CompanyID: 1,
		InvoiceID: 2,
		Type:      "product",
		Qty:       1,
		Rate:      10.00,
		ItemID:    123,
		Discount:  0.00,
		Amount:    10.00,
	}
	invoiceItem, err := store.CreateInvoiceItem(param)
	return store, invoiceItem, err
}
