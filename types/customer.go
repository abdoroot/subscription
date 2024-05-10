package types

import "time"

type Customer struct {
	ID          int       `db:"id"`
	Type        string    `db:"type"`
	DisplayName string    `db:"display_name"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	CompanyName string    `db:"company_name"`
	Email       string    `db:"email"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (c Customer) Validate() bool {
	return true
}
