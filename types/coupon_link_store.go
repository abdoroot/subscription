package types

import "time"

type CouponLink struct {
	ID       int    `db:"id"`
	CouponId int    `db:"coupon_id"`
	LinkType string `db:"link_type"`
	LinkToId int    `db:"link_to_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c CouponLink) Validate() bool {
	return true
}
