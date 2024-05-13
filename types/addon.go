package types

import (
	"time"
)

type Addon struct {
	ID              int       `db:"id"`
	CompanyID       int       `db:"company_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	DisplayInWidget bool      `db:"display_in_widget"`
	PricingModelID  int       `db:"pricing_model_id"`
	UnitID          int       `db:"unit_id"`
	BillingType     string    `db:"billing_type"`
	BillingCycle    string    `db:"billing_cycle"`
	Price           float64   `db:"price"`
	LinkedToPlans   string    `db:"linked_to_plans"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func (c Addon) Validate() bool {
	return true
}
