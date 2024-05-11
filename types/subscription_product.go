package types

import "time"

type SubscriptionProduct struct {
	Id          int       `db:"id"`
	CompanyId   int       `db:"company_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
