package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

var insertedProduct types.Product

func TestCreateProduct(t *testing.T) {
	param := types.Product{
		CompanyId:         1,
		Name:              "water",
		Type:              "product",
		UnitId:            1,
		Price:             1.55,
		ImageAttachmentId: 1,
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Errorf("error connection to db %v", err)
	}
	ps := NewProductStore(db)
	p, err := ps.CreateProduct(param)
	if err != nil {
		t.Errorf("error CreateProduct %v", err)
	}
	insertedProduct = p

	assert.Greater(t, p.Id, 0)
	assert.Greater(t, p.CreatedAt, time.Time{})
}

func TestUpdateProductById(t *testing.T) {
	if insertedProduct.Id == 0 {
		t.Error("TestUpdateProductById depend on insertedProduct")
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Errorf("error connection to db %v", err)
	}
	ps := NewProductStore(db)
	param := types.UpdateProductRequest{
		Name:              "coca cola",
		Type:              "product",
		UnitId:            1,
		Price:             3.5,
		ImageAttachmentId: 1,
	}
	err = ps.UpdateProductById(param, insertedProduct.Id)
	if err != nil {
		t.Errorf("error connection to db %v", err)
	}
}

func TestDeleteProductByID(t *testing.T) {
	if insertedProduct.Id == 0 {
		t.Error("TestUpdateProductById depend on insertedProduct")
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Errorf("error connection to db %v", err)
	}
	ps := NewProductStore(db)

	if err := ps.DeleteProductByID(insertedProduct.Id); err != nil {
		t.Errorf("error DeleteProductByID %v", err)
	}
}
