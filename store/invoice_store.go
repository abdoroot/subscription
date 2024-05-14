package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type InvoiceStore struct {
	db         *sqlx.DB
	itemsStore *invoiceItemStore
}

func NewInvoiceStore(db *sqlx.DB, itemsStore *invoiceItemStore) *InvoiceStore {
	return &InvoiceStore{db: db, itemsStore: itemsStore}
}

func (s *InvoiceStore) CreateInvoice(param types.Invoice) (types.Invoice, error) {

	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `INSERT INTO invoices(company_id, customer_id, order_number, invoice_date,due_date, subject, terms_and_conditions, attachments, created_at, updated_at)
			   VALUES (:company_id, :customer_id, :order_number, :invoice_date,:due_date, :subject, :terms_and_conditions, :attachments, :created_at, :updated_at)
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

	//insert invoice items
	if len(param.Items) > 0 {
		for _, item := range param.Items {
			item.InvoiceID = param.ID
			if _, err = s.itemsStore.CreateInvoiceItem(item); err != nil {
				//delete the invoice
				s.DeleteInvoiceByID(param.ID)
				return types.Invoice{}, err
			}
		}
	}

	return param, nil
}

func (s *InvoiceStore) UpdateInvoiceByID(param types.Invoice, id int) (err error) {
	//Update the invoice
	param.UpdatedAt = time.Now()
	param.ID = id
	query := `UPDATE invoices
			  SET customer_id = :customer_id,
			      order_number = :order_number,
			      invoice_date = :invoice_date,
			      due_date = :due_date,
			      subject = :subject,
			      terms_and_conditions = :terms_and_conditions,
			      attachments = :attachments,
			      updated_at = :updated_at
			  WHERE id = :id`
	if _, err := s.db.NamedExec(query, param); err != nil {
		return err
	}
	// Update invoice items
	if len(param.Items) > 0 {
		for _, item := range param.Items {
			if item.ID > 0 {
				//update the item
				if err := s.itemsStore.UpdateInvoiceItemByID(item, item.ID); err != nil {
					logrus.Errorf("error UpdateInvoiceItemByID item %v", item)
					return err
				}
			} else if item.ID == 0 {
				if _, err := s.itemsStore.CreateInvoiceItem(item); err != nil {
					logrus.Errorf("error CreateInvoiceItem form UpdateInvoiceByID func item %v", item)
					return err
				}
			}
		}
	}

	return nil
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
	//First delete invoice items
	if err := s.itemsStore.DeleteInvoiceItemByInvoiceID(id); err != nil {
		return err
	}
	//Second delete the invoice
	query := `DELETE FROM invoices WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
