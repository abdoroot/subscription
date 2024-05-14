package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type CustomerSubscriptionStore struct {
	db *sqlx.DB
}

func NewCustomerSubscriptionStore(db *sqlx.DB) *CustomerSubscriptionStore {
	return &CustomerSubscriptionStore{db: db}
}

func (s *CustomerSubscriptionStore) CreateCustomerSubscription(subscription types.CustomerSubscription) (types.CustomerSubscription, error) {
	subscription.CreatedAt = time.Now()
	subscription.UpdatedAt = time.Now()
	query := `INSERT INTO customer_subscriptions(company_id, customer_id, plan_id, start_date, never_expire, expire_after_cycle, qty, price, amount, reference, offline_payment, created_at, updated_at)
               VALUES (:company_id, :customer_id, :plan_id, :start_date, :never_expire, :expire_after_cycle, :qty, :price, :amount, :reference, :offline_payment, :created_at, :updated_at)
               RETURNING id`
	row, err := s.db.NamedQuery(query, subscription)
	if err != nil {
		return types.CustomerSubscription{}, err
	}
	if row.Next() {
		if err := row.Scan(&subscription.ID); err != nil {
			return types.CustomerSubscription{}, err
		}
	}
	return subscription, nil
}

func (s *CustomerSubscriptionStore) UpdateCustomerSubscriptionByID(subscription types.CustomerSubscription, id int) error {
	subscription.UpdatedAt = time.Now()
	subscription.ID = id
	query := `UPDATE customer_subscriptions
              SET customer_id = :customer_id,
                  plan_id = :plan_id,
                  start_date = :start_date,
                  never_expire = :never_expire,
                  expire_after_cycle = :expire_after_cycle,
                  qty = :qty,
                  price = :price,
                  amount = :amount,
                  reference = :reference,
                  offline_payment = :offline_payment,
                  updated_at = :updated_at
              WHERE id = :id`
	_, err := s.db.NamedExec(query, subscription)
	return err
}

func (s *CustomerSubscriptionStore) GetCustomerSubscriptionByID(id int) (types.CustomerSubscription, error) {
	var subscription types.CustomerSubscription
	query := `SELECT * FROM customer_subscriptions WHERE id = $1`
	err := s.db.Get(&subscription, query, id)
	return subscription, err
}

func (s *CustomerSubscriptionStore) GetCustomerSubscriptionsByCustomerID(customerID int) ([]types.CustomerSubscription, error) {
	var subscriptions []types.CustomerSubscription
	query := `SELECT * FROM customer_subscriptions WHERE customer_id = $1`
	err := s.db.Select(&subscriptions, query, customerID)
	return subscriptions, err
}

func (s *CustomerSubscriptionStore) GetCustomerSubscriptionsByCompanyID(companyID int) ([]types.CustomerSubscription, error) {
	var subscriptions []types.CustomerSubscription
	query := `SELECT * FROM customer_subscriptions WHERE company_id = $1`
	err := s.db.Select(&subscriptions, query, companyID)
	return subscriptions, err
}

func (s *CustomerSubscriptionStore) GetAllCustomerSubscriptions() ([]types.CustomerSubscription, error) {
	var subscriptions []types.CustomerSubscription
	query := `SELECT * FROM customer_subscriptions`
	err := s.db.Select(&subscriptions, query)
	return subscriptions, err
}

func (s *CustomerSubscriptionStore) DeleteCustomerSubscriptionByID(id int) error {
	query := `DELETE FROM customer_subscriptions WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
