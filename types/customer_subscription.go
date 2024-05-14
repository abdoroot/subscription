package types

import "time"

type CustomerSubscription struct {
	ID               int       `db:"id"`
	CompanyID        int       `db:"company_id"`
	CustomerID       int       `db:"customer_id"`
	PlanID           int       `db:"plan_id"`
	StartDate        time.Time `db:"start_date"`
	NeverExpire      bool      `db:"never_expire"`
	ExpireAfterCycle int       `db:"expire_after_cycle"`
	Qty              int       `db:"qty"`
	Price            float64   `db:"price"`
	Amount           float64   `db:"amount"`
	Reference        string    `db:"reference"`
	OfflinePayment   bool      `db:"offline_payment"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
