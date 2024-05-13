package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type couponStore struct {
	db *sqlx.DB
}

func NewCouponStore(db *sqlx.DB) *couponStore {
	return &couponStore{
		db: db,
	}
}

func (s *couponStore) CreateCoupon(param types.Coupon) (types.Coupon, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `INSERT INTO coupons(company_id, name, discount_value, discount_type, redemption_type, max_redemptions, redemption_count, linked_to_plans, linked_to_addons, expire_date, created_at, updated_at)
			   VALUES (:company_id, :name, :discount_value, :discount_type, :redemption_type, :max_redemptions, :redemption_count, :linked_to_plans, :linked_to_addons, :expire_date, :created_at, :updated_at)
			   RETURNING id`

	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Coupon{}, err
	}
	if row.Next() {
		if err := row.Scan(&param.ID); err != nil {
			return types.Coupon{}, err
		}
	}
	return param, nil
}

func (s *couponStore) UpdateCouponById(param types.Coupon, id int) error {
	param.UpdatedAt = time.Now()
	param.ID = id
	query := `UPDATE coupons
			  SET name = :name, 
			      discount_value = :discount_value, 
			      discount_type = :discount_type, 
			      redemption_type = :redemption_type, 
			      max_redemptions = :max_redemptions, 
			      redemption_count = :redemption_count, 
			      linked_to_plans = :linked_to_plans, 
			      linked_to_addons = :linked_to_addons, 
			      expire_date = :expire_date, 
			      updated_at = :updated_at
			  WHERE id = :id`

	_, err := s.db.NamedExec(query, param)
	return err
}

func (s *couponStore) GetCouponById(id int) (types.Coupon, error) {
	var coupon types.Coupon
	query := `SELECT * FROM coupons WHERE id = $1`
	err := s.db.Get(&coupon, query, id)
	return coupon, err
}

func (s *couponStore) GetCouponByCompanyId(id int) ([]types.Coupon, error) {
	var coupons []types.Coupon
	query := `SELECT * FROM coupons WHERE company_id = $1`
	err := s.db.Select(&coupons, query, id)
	return coupons, err
}

func (s *couponStore) GetAll() ([]types.Coupon, error) {
	var coupons []types.Coupon
	query := `SELECT * FROM coupons`
	err := s.db.Select(&coupons, query)
	return coupons, err
}

func (s *couponStore) DeleteCouponById(id int) error {
	query := `DELETE FROM coupons WHERE id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	return err
}
