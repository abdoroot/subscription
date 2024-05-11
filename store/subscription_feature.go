package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type subscriptionFeatureStore struct {
	db *sqlx.DB
}

func NewSubscriptionFeatureStore(db *sqlx.DB) *subscriptionFeatureStore {
	return &subscriptionFeatureStore{
		db: db,
	}
}

func (s *subscriptionFeatureStore) CreateFeature(param types.SubscriptionFeature) (types.SubscriptionFeature, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `insert into subscription_features (subscription_plan_id,name,created_at,updated_at) 
	values(:subscription_plan_id,:name,:created_at,:updated_at) returning id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.SubscriptionFeature{}, err
	}
	var lastInsertedId int
	if row.Next() {
		if err := row.Scan(&lastInsertedId); err != nil {
			return types.SubscriptionFeature{}, err
		}
	}
	param.Id = lastInsertedId
	return param, nil
}

func (s *subscriptionFeatureStore) UpdateFeatureById(param types.SubscriptionFeature, id int) error {
	param.UpdatedAt = time.Now()
	param.Id = id
	query := `update subscription_features set name= :name,updated_at= :updated_at where id = :id  `
	_, err := s.db.NamedQuery(query, param)
	if err != nil {
		return err
	}
	return nil
}

func (s *subscriptionFeatureStore) GetAll() ([]types.SubscriptionFeature, error) {
	var p []types.SubscriptionFeature
	query := `select * from subscription_features`
	if err := s.db.Select(&p, query); err != nil {
		return []types.SubscriptionFeature{}, err
	}
	return p, nil
}

func (s *subscriptionFeatureStore) GetFeaturesByCompanyId(id int) ([]types.SubscriptionFeature, error) {
	var p []types.SubscriptionFeature
	query := `select * from subscription_features where company_id =$1`
	if err := s.db.Select(&p, query, id); err != nil {
		return []types.SubscriptionFeature{}, err
	}
	return p, nil
}

func (s *subscriptionFeatureStore) GetFeatureById(id int) (types.SubscriptionFeature, error) {
	var p types.SubscriptionFeature
	query := `select * from subscription_features where id =$1`
	if err := s.db.Get(&p, query, id); err != nil {
		return types.SubscriptionFeature{}, err
	}
	return p, nil
}

func (s *subscriptionFeatureStore) DeleteFeatureById(id int) error {
	query := `delete from subscription_features where id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
