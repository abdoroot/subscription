package types

import (
	"time"
)

type Coupon struct {
    ID              int       `db:"id"`
    CompanyID       int       `db:"company_id"`
    Name            string    `db:"name"`
    DiscountValue   float64   `db:"discount_value"`
    DiscountType    string    `db:"discount_type"`
    RedemptionType  string    `db:"redemption_type"`
    MaxRedemptions  int       `db:"max_redemptions"`
    RedemptionCount int       `db:"redemption_count"`
    LinkedToPlans   string    `db:"linked_to_plans"`
    LinkedToAddons  string    `db:"linked_to_addons"`
    ExpireDate      time.Time `db:"expire_date"`
    CreatedAt       time.Time `db:"created_at"`
    UpdatedAt       time.Time `db:"updated_at"`
}

