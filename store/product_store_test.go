package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	ps, p, err := createProduct()
	defer ps.DeleteProductByID(p.Id)
	if err != nil {
		t.Errorf("error createProduct %v", err)
	}
	assert.Greater(t, p.Id, 0)
	assert.Greater(t, p.CreatedAt, time.Time{})
}

func TestUpdateProductById(t *testing.T) {
	ps, p, err := createProduct()
	defer ps.DeleteProductByID(p.Id)
	if err != nil {
		t.Errorf("error createProduc %v", err)
	}
	param := types.UpdateProductRequest{
		Name:              "coca cola",
		Type:              "product",
		UnitId:            1,
		Price:             3.5,
		ImageAttachmentId: 1,
	}
	err = ps.UpdateProductById(param, p.Id)
	if err != nil {
		t.Errorf("error connection to db %v", err)
	}
}

func TestDeleteProductByID(t *testing.T) {
	ps, p, err := createProduct()
	defer ps.DeleteProductByID(p.Id)
	if err != nil {
		t.Errorf("error createProduc %v", err)
	}
	if err := ps.DeleteProductByID(p.Id); err != nil {
		t.Errorf("error DeleteProductByID %v", err)
	}
}

func TestGetProducts(t *testing.T) {
	ps, p, err := createProduct()
	defer ps.DeleteProductByID(p.Id)
	if err != nil {
		t.Errorf("error createProduc %v", err)
	}

	t.Run("get all", func(t *testing.T) {
		products, err := ps.GetAll()
		if err != nil {
			t.Errorf("error GetAll %v", err)
		}

		assert.Greater(t, len(products), 0)
	})

	t.Run("get by company id", func(t *testing.T) {
		products, err := ps.GetAllByCompanyId(p.CompanyId)
		if err != nil {
			t.Errorf("error GetAll %v", err)
		}

		assert.Greater(t, len(products), 0)
	})
}

// help testing func
func createProduct() (*productStore, types.Product, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.Product{}, err
	}

	param := types.Product{
		CompanyId:         1,
		Name:              "water",
		Type:              "product",
		UnitId:            1,
		Price:             1.55,
		ImageAttachmentId: 1,
	}

	ps := NewProductStore(db)
	p, err := ps.CreateProduct(param)
	if err != nil {
		return nil, types.Product{}, err
	}
	return ps, p, nil
}
