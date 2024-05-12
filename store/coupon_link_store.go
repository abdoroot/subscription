package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type couponLinkStore struct {
	db *sqlx.DB
}

func NewCouponLinkStore(db *sqlx.DB) *couponLinkStore {
	return &couponLinkStore{
		db: db,
	}
}

func (s *couponLinkStore) CreateCouponLink(param types.CouponLink) (types.CouponLink, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `insert into coupon_links(coupon_id,link_type,link_to_id,created_at,updated_at) 
	values(:coupon_id,:link_type,:link_to_id,:created_at,:updated_at) returning id
	`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.CouponLink{}, err
	}

	if row.Next() {
		if err := row.Scan(&param.ID); err != nil {
			return types.CouponLink{}, err
		}
	}
	return param, err
}

func (s *couponLinkStore) GetCouponLinkById(id int) (types.CouponLink, error) {
	var link types.CouponLink
	query := `select * from coupon_links  where id= $1`
	if err := s.db.Get(&link, query,id); err != nil {
		return types.CouponLink{}, err
	}
	return link, nil
}

func (s *couponLinkStore) DeleteCouponLinkById(id int) error {
	query := `DELETE FROM coupon_links WHERE id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}
