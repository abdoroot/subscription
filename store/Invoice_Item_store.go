package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type invoiceItemStore struct {
	db *sqlx.DB
}

func NewInvoiceItemStore(db *sqlx.DB) *invoiceItemStore {
	return &invoiceItemStore{db: db}
}

func (s *invoiceItemStore) CreateInvoiceItem(param types.InvoiceItem) (types.InvoiceItem, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `INSERT INTO invoice_items(company_id, invoice_id, item_id, type, qty, rate, discount, amount, created_at, updated_at)
               VALUES (:company_id, :invoice_id, :item_id, :type, :qty, :rate, :discount, :amount, :created_at, :updated_at)
               RETURNING id`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.InvoiceItem{}, err
	}
	if row.Next() {
		if err := row.Scan(&param.ID); err != nil {
			return types.InvoiceItem{}, err
		}
	}
	return param, nil
}

func (s *invoiceItemStore) UpdateInvoiceItemByID(param types.InvoiceItem, id int) error {
	param.UpdatedAt = time.Now()
	param.ID = id
	query := `UPDATE invoice_items
              SET company_id = :company_id,
                  invoice_id = :invoice_id,
                  item_id = :item_id,
                  type = :type,
                  qty = :qty,
                  rate = :rate,
                  discount = :discount,
                  amount = :amount,
                  updated_at = :updated_at
              WHERE id = :id`
	_, err := s.db.NamedExec(query, param)
	return err
}

func (s *invoiceItemStore) GetInvoiceItemByID(id int) (types.InvoiceItem, error) {
	var invoiceItem types.InvoiceItem
	query := `SELECT * FROM invoice_items WHERE id = $1`
	err := s.db.Get(&invoiceItem, query, id)
	return invoiceItem, err
}

func (s *invoiceItemStore) GetInvoiceItemsByCompanyID(companyID int) ([]types.InvoiceItem, error) {
	var invoiceItems []types.InvoiceItem
	query := `SELECT * FROM invoice_items WHERE company_id = $1`
	err := s.db.Select(&invoiceItems, query, companyID)
	return invoiceItems, err
}

func (s *invoiceItemStore) GetAllInvoiceItems() ([]types.InvoiceItem, error) {
	var invoiceItems []types.InvoiceItem
	query := `SELECT * FROM invoice_items`
	err := s.db.Select(&invoiceItems, query)
	return invoiceItems, err
}

func (s *invoiceItemStore) DeleteInvoiceItemByID(id int) error {
	query := `DELETE FROM invoice_items WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *invoiceItemStore) DeleteInvoiceItemByInvoiceID(id int) error {
	query := `DELETE FROM invoice_items WHERE invoice_id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
