package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type PaymentStore struct {
	db *sqlx.DB
}

func NewPaymentStore(db *sqlx.DB) *PaymentStore {
	return &PaymentStore{db: db}
}

func (s *PaymentStore) CreatePayment(payment types.Payment) (types.Payment, error) {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	query := `INSERT INTO payments(company_id, customer_id, invoice_id, amount_received, payment_date, payment_method_id, reference, attachments, created_at, updated_at)
               VALUES (:company_id, :customer_id, :invoice_id, :amount_received, :payment_date, :payment_method_id, :reference, :attachments, :created_at, :updated_at)
               RETURNING id`
	row, err := s.db.NamedQuery(query, payment)
	if err != nil {
		return types.Payment{}, err
	}
	if row.Next() {
		if err := row.Scan(&payment.ID); err != nil {
			return types.Payment{}, err
		}
	}
	return payment, nil
}

func (s *PaymentStore) UpdatePaymentByID(payment types.Payment, id int) error {
	payment.UpdatedAt = time.Now()
	payment.ID = id
	query := `UPDATE payments
              SET customer_id = :customer_id,
                  invoice_id = :invoice_id,
                  amount_received = :amount_received,
                  payment_date = :payment_date,
                  payment_method_id = :payment_method_id,
                  reference = :reference,
                  attachments = :attachments,
                  updated_at = :updated_at
              WHERE id = :id`
	_, err := s.db.NamedExec(query, payment)
	return err
}

func (s *PaymentStore) GetPaymentByID(id int) (types.Payment, error) {
	var payment types.Payment
	query := `SELECT * FROM payments WHERE id = $1`
	err := s.db.Get(&payment, query, id)
	return payment, err
}

func (s *PaymentStore) GetPaymentsByInvoiceID(invoiceID int) ([]types.Payment, error) {
	var payments []types.Payment
	query := `SELECT * FROM payments WHERE invoice_id = $1`
	err := s.db.Select(&payments, query, invoiceID)
	return payments, err
}

func (s *PaymentStore) GetPaymentsByCompanyID(companyID int) ([]types.Payment, error) {
	var payments []types.Payment
	query := `SELECT * FROM payments WHERE company_id = $1`
	err := s.db.Select(&payments, query, companyID)
	return payments, err
}

func (s *PaymentStore) GetPaymentsByCustomerID(customerID int) ([]types.Payment, error) {
	var payments []types.Payment
	query := `SELECT * FROM payments WHERE customer_id = $1`
	err := s.db.Select(&payments, query, customerID)
	return payments, err
}

func (s *PaymentStore) GetAllPayments() ([]types.Payment, error) {
	var payments []types.Payment
	query := `SELECT * FROM payments`
	err := s.db.Select(&payments, query)
	return payments, err
}

func (s *PaymentStore) DeletePaymentByID(id int) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
