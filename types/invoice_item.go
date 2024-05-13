package types

import "time"

type InvoiceItem struct {
	ID        int       `db:"id"`
	CompanyID int       `db:"company_id"`
	InvoiceID int       `db:"invoice_id"`
	ItemID    int       `db:"item_id"`
	Type      string    `db:"type"`
	Qty       int       `db:"qty"`
	Rate      float64   `db:"rate"`
	Discount  float64   `db:"discount"`
	Amount    float64   `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
