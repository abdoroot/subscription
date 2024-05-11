package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type subscriptionProductsStore struct {
	db *sqlx.DB
}

func NewSubscriptionProductsStore(db *sqlx.DB) *subscriptionProductsStore {
	return &subscriptionProductsStore{
		db: db,
	}
}

func (s *subscriptionProductsStore) CreateProduct(param types.SubscriptionProduct) (types.SubscriptionProduct, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `insert into subscription_products (company_id,name,description,created_at,updated_at) 
	values(:company_id,:name,:description,:created_at,:updated_at) returning id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.SubscriptionProduct{}, err
	}
	var lastInsertedId int
	if row.Next() {
		if err := row.Scan(&lastInsertedId); err != nil {
			return types.SubscriptionProduct{}, err
		}
	}
	param.Id = lastInsertedId
	return param, nil
}

func (s *subscriptionProductsStore) UpdateProductById(param types.SubscriptionProduct, id int) error {
	param.UpdatedAt = time.Now()
	query := `update subscription_products set name= :name,description= :description,updated_at= :updated_at;`
	_, err := s.db.NamedQuery(query, param)
	if err != nil {
		return err
	}
	return nil
}

func (s *subscriptionProductsStore) DeleteProductById(id int) error {
	query := `delete from subscription_products where id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
