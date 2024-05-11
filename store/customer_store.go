package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type customerStore struct {
	db       *sqlx.DB
	addStore *addressStore
}

func NewCustomerStore(db *sqlx.DB, addStore *addressStore) *customerStore {
	return &customerStore{
		db:       db,
		addStore: addStore,
	}
}

func (s *customerStore) CreateCustomer(customer types.Customer) (types.Customer, error) {
	customer.CreatedAt = time.Now().UTC()
	customer.UpdatedAt = time.Now().UTC()
	tx := s.db.MustBegin()
	query := `
		INSERT INTO customers (type,company_id, display_name, first_name, last_name, company_name, email, created_at, updated_at)
		VALUES (:type,:company_id ,:display_name, :first_name, :last_name, :company_name, :email, :created_at, :updated_at)
		RETURNING id;
	`
	rows, err := s.db.NamedQuery(query, customer)

	if err != nil {
		tx.Rollback()
		return types.Customer{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&customer.ID); err != nil {
			tx.Rollback()
			return types.Customer{}, err
		}
	}

	if len(customer.Address) > 0 {
		addressInputs := customer.Address
		//clear the customer.Address to use later
		customer.Address = customer.Address[:0]
		for _, v := range addressInputs {
			req := v
			req.CustomerID = customer.ID
			addr, err := s.addStore.CreateAddress(req)
			if err != nil {
				tx.Rollback()
				return types.Customer{}, err
			}
			customer.Address = append(customer.Address, addr)
		}
	}

	tx.Commit()
	return customer, nil
}

func (s *customerStore) UpdateCustomer(customer types.Customer) error {
	customer.UpdatedAt = time.Now().UTC()
	query := `
		UPDATE customers
		SET type = :type, display_name = :display_name, first_name = :first_name, last_name = :last_name, company_name = :company_name, email = :email, updated_at = :updated_at
		WHERE id = :id;
	`
	_, err := s.db.NamedExec(query, customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *customerStore) DeleteCustomerByID(id int) error {
	//First delete customer address
	err := s.addStore.DeleteAddressByCustomerId(id)
	if err != nil {
		return err
	}
	//Second Delete the customer
	query := `DELETE FROM customers WHERE id = :id;`
	_, err = s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}

	return nil
}

func (s *customerStore) GetCustomerByID(id int) (types.Customer, error) {
	var customer types.Customer
	query := `SELECT * FROM customers WHERE id = $1;`
	err := s.db.Get(&customer, query, id)
	if err != nil {
		return types.Customer{}, err
	}
	return customer, nil
}

func (s *customerStore) GetCustomersByCompanyId(id int) ([]types.Customer, error) {
	var customers []types.Customer
	query := `SELECT * FROM customers WHERE company_id = $1;`
	err := s.db.Select(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *customerStore) GetAll() ([]types.Customer, error) {
	var customers []types.Customer
	query := `SELECT * FROM customers`
	err := s.db.Select(&customers, query)
	if err != nil {
		return []types.Customer{}, err
	}
	return customers, nil
}
