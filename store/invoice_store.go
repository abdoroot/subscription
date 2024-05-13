package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type InvoiceStore struct {
	db *sqlx.DB
}

func NewInvoiceStore(db *sqlx.DB) *InvoiceStore {
	return &InvoiceStore{db: db}
}

func (s *InvoiceStore) CreateInvoice(param types.Invoice) (types.Invoice, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `INSERT INTO invoices(company_id, customer_id, order_number, invoice_date, subject, terms_and_conditions, attachments, created_at, updated_at)
			   VALUES (:company_id, :customer_id, :order_number, :invoice_date, :subject, :terms_and_conditions, :attachments, :created_at, :updated_at)
			   RETURNING id`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Invoice{}, err
	}
	if row.Next() {
		if err := row.Scan(&param.ID); err != nil {
			return types.Invoice{}, err
		}
	}
	return param, nil
}

func (s *InvoiceStore) UpdateInvoiceByID(param types.Invoice, id int) error {
	param.UpdatedAt = time.Now()
	param.ID = id
	query := `UPDATE invoices
			  SET customer_id = :customer_id,
			      order_number = :order_number,
			      invoice_date = :invoice_date,
			      subject = :subject,
			      terms_and_conditions = :terms_and_conditions,
			      attachments = :attachments,
			      updated_at = :updated_at
			  WHERE id = :id`
	_, err := s.db.NamedExec(query, param)
	return err
}

func (s *InvoiceStore) GetInvoiceByID(id int) (types.Invoice, error) {
	var invoice types.Invoice
	query := `SELECT * FROM invoices WHERE id = $1`
	err := s.db.Get(&invoice, query, id)
	return invoice, err
}

func (s *InvoiceStore) GetInvoicesByCompanyID(companyID int) ([]types.Invoice, error) {
	var invoices []types.Invoice
	query := `SELECT * FROM invoices WHERE company_id = $1`
	err := s.db.Select(&invoices, query, companyID)
	return invoices, err
}

func (s *InvoiceStore) GetAllInvoices() ([]types.Invoice, error) {
	var invoices []types.Invoice
	query := `SELECT * FROM invoices`
	err := s.db.Select(&invoices, query)
	return invoices, err
}

func (s *InvoiceStore) DeleteInvoiceByID(id int) error {
	query := `DELETE FROM invoices WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
