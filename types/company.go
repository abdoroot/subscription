package types

import "time"

type Company struct {
	Id                    int       `db:"id"`
	Name                  string    `db:"name"`
	CountryId             int       `db:"country_id"`
	CityId                int       `db:"city_id"`
	IsVatRegistered       bool      `db:"is_vat_registered"`
	TaxRegistrationNumber string    `db:"tax_registration_number"`
	VatRegistrationNumber string    `db:"vat_registration_number"`
	CurrencyId            int       `db:"currency_id"`
	BillingType           string    `db:"billing_type"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}

func (c Company) Validate() bool {
	return true
}
