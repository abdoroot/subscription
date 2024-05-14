package types

import (
	"time"

	"github.com/lib/pq"
)

type Invoice struct {
	ID                 int           `db:"id"`
	CompanyID          int           `db:"company_id"`
	CustomerID         int           `db:"customer_id"`
	OrderNumber        string        `db:"order_number"`
	InvoiceDate        time.Time     `db:"invoice_date"`
	DueDate            time.Time     `db:"due_date"`
	Subject            string        `db:"subject"`
	TermsAndConditions string        `db:"terms_and_conditions"`
	Attachments        pq.Int32Array `db:"attachments"`
	Items              []InvoiceItem `json:"invoice_items"`
	CreatedAt          time.Time     `db:"created_at"`
	UpdatedAt          time.Time     `db:"updated_at"`
}
