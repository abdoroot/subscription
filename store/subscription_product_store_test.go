package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubProduct(t *testing.T) {
	sps, p, err := createSubProduct()
	_ = sps
	if err != nil {
		t.Error("error creating createSubProduct ", err)
	}
	defer sps.DeleteProductById(p.Id)
	assert.Greater(t, p.Id, 0)
}

func TestUpdateProduct(t *testing.T) {
	sps, p, err := createSubProduct()
	if err != nil {
		t.Error("error creating createSubProduct ", err)
	}
	defer sps.DeleteProductById(p.Id)
	param := types.SubscriptionProduct{
		Name:        "update milk",
		Description: "update milk description",
	}
	if err := sps.UpdateProductById(param, p.Id); err != nil {
		t.Error("error UpdateProductById ", err)
	}
	assert.Greater(t, p.Id, 0)
}

func TestDeleteProduct(t *testing.T) {
	sps, p, err := createSubProduct()
	if err != nil {
		t.Error("error creating createSubProduct ", err)
	}
	defer sps.DeleteProductById(p.Id)
	if err := sps.DeleteProductById(p.Id); err != nil {
		t.Error("error UpdateProductById ", err)
	}
	assert.Greater(t, p.Id, 0)
}

func TestGetProduct(t *testing.T) {
	sps, p, err := createSubProduct()
	if err != nil {
		t.Error("error creating createSubProduct ", err)
	}
	defer sps.DeleteProductById(p.Id)
	t.Run("GetProductById", func(t *testing.T) {
		product, err := sps.GetProductById(p.Id)
		if err != nil {
			t.Error("error GetProductById ", err)
		}
		assert.Greater(t, product.Id, 0)
	})

	t.Run("GetProductsByCompanyId", func(t *testing.T) {
		products, err := sps.GetProductsByCompanyId(p.CompanyId)
		if err != nil {
			t.Error("error GetProductsByCompanyId ", err)
		}
		assert.Greater(t, len(products), 0)
	})

	t.Run("GetAll", func(t *testing.T) {
		products, err := sps.GetAll()
		if err != nil {
			t.Error("error GetAll", err)
		}
		assert.Greater(t, len(products), 0)
	})
}

// test helper func
func createSubProduct() (*subscriptionProductsStore, types.SubscriptionProduct, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.SubscriptionProduct{}, err
	}
	sps := NewSubscriptionProductsStore(db)
	param := types.SubscriptionProduct{
		CompanyId:   5,
		Name:        "milk",
		Description: "milk description",
	}
	p, err := sps.CreateProduct(param)
	if err != nil {
		return nil, types.SubscriptionProduct{}, err
	}
	return sps, p, nil
}
