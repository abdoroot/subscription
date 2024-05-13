package types

import "time"

type Invoice struct {
	ID                 int       `db:"id"`
	CompanyID          int       `db:"company_id"`
	CustomerID         int       `db:"customer_id"`
	OrderNumber        string    `db:"order_number"`
	InvoiceDate        time.Time `db:"invoice_date"`
	Subject            string    `db:"subject"`
	TermsAndConditions string    `db:"terms_and_conditions"`
	Attachments        []int     `db:"attachments"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
