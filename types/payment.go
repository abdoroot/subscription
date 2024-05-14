package types

import (
	"time"

	"github.com/lib/pq"
)

type Payment struct {
	ID              int           `db:"id"`
	CompanyID       int           `db:"company_id"`
	CustomerID      int           `db:"customer_id"`
	InvoiceID       int           `db:"invoice_id"`
	AmountReceived  float64       `db:"amount_received"`
	PaymentDate     time.Time     `db:"payment_date"`
	PaymentMethodID int           `db:"payment_method_id"`
	Reference       string        `db:"reference"`
	Attachments     pq.Int32Array `db:"attachments"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
}
