package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type productStore struct {
	db *sqlx.DB
}

func NewProductStore(db *sqlx.DB) *productStore {
	return &productStore{
		db: db,
	}
}

func (s *productStore) CreateProduct(param types.Product) (types.Product, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `insert into products (company_id,name,type,unit_id,price,description,image_attachment_id,created_at,updated_at) 
	values(:company_id,:name,:type,:unit_id,:price,:description,:image_attachment_id,:created_at,:updated_at) returning id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Product{}, err
	}
	var lastInsertedId int
	if row.Next() {
		if err := row.Scan(&lastInsertedId); err != nil {
			return types.Product{}, err
		}
	}
	param.Id = lastInsertedId
	return param, nil
}

func (s *productStore) UpdateProductById(param types.UpdateProductRequest, id int) error {
	param.Id = id
	param.UpdatedAt = time.Now()
	query := `update products set name=:name,type=:type,unit_id=:unit_id,price=:price,description=:description,image_attachment_id=:image_attachment_id where id=:id`
	_, err := s.db.NamedExec(query, param)
	if err != nil {
		return err
	}
	return nil
}

func (s *productStore) DeleteProductByID(id int) error {
	query := `DELETE FROM products WHERE id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (s *productStore) DeleteProductByCompanyID(id int) error {
	query := `DELETE FROM products WHERE company_id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}
