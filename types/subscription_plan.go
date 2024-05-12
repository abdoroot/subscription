package types

import (
	"time"
)

type SubscriptionPlan struct {
	Id                int     `db:"id"`
	CompanyId         int     `db:"company_id"`
	ProductId         int     `db:"product_id"`
	Name              string  `db:"name"`
	Code              string  `db:"code"`
	UnitId            int     `db:"unit_id"`
	IsActive          bool    `db:"is_active"`
	Price             float32 `db:"price"`
	Type              string  `db:"type"`
	BillingEvery      string  `db:"billing_every"`
	BillingEveryCount int     `db:"billing_every_count"`
	BillingCycle      string  `db:"billing_cycle"`
	BillingCycleCount int     `db:"billing_cycle_count"`
	TrialPeriod       int     `db:"trial_period"`
	SetupFee          float32 `db:"setup_fee"`
	HaveFeatures      bool    `db:"have_features"`
	Description       string  `db:"description"`
	Features          []SubscriptionFeature
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
