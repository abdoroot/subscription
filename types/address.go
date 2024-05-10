package types

import "time"

type Address struct {
	ID         int       `db:"id"`
	Type       string    `db:"type"` // Should be of type address_type_enums, handle it in validation
	CustomerID int       `db:"customer_id"`
	CountryID  int       `db:"country_id"`
	CityID     int       `db:"city_id"`
	Line1      string    `db:"line1"`
	Line2      string    `db:"line2"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
