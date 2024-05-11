package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type roleStore struct {
	db *sqlx.DB
}

func NewRoleStore(db *sqlx.DB) *roleStore {
	return &roleStore{
		db: db,
	}
}

func (s *roleStore) CreateRole(role types.Role) (types.Role, error) {
	role.CreatedAt = time.Now().UTC()
	role.UpdatedAt = time.Now().UTC()
	query := `
        INSERT INTO roles (company_id, name, can_add_products, can_update_products, can_delete_products, can_view_products, can_add_customers, can_update_customers, can_delete_customers, can_view_customers, can_add_subscriptions, can_update_subscriptions, can_delete_subscriptions, can_view_subscriptions, can_add_items, can_update_items, can_delete_items, can_view_items, can_add_invoices, can_update_invoices, can_delete_invoices, can_view_invoices, can_add_payments, can_update_payments, can_delete_payments, can_view_payments, can_view_dashboard, created_at, updated_at)
        VALUES (:company_id, :name, :can_add_products, :can_update_products, :can_delete_products, :can_view_products, :can_add_customers, :can_update_customers, :can_delete_customers, :can_view_customers, :can_add_subscriptions, :can_update_subscriptions, :can_delete_subscriptions, :can_view_subscriptions, :can_add_items, :can_update_items, :can_delete_items, :can_view_items, :can_add_invoices, :can_update_invoices, :can_delete_invoices, :can_view_invoices, :can_add_payments, :can_update_payments, :can_delete_payments, :can_view_payments, :can_view_dashboard, :created_at, :updated_at)
        RETURNING id;
    `
	rows, err := s.db.NamedQuery(query, role)
	if err != nil {
		return types.Role{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&role.ID); err != nil {
			return types.Role{}, err
		}
	}
	return role, nil
}

func (s *roleStore) UpdateRole(role types.Role) error {
	role.UpdatedAt = time.Now().UTC()
	query := `
        UPDATE roles
        SET name = :name, can_add_products = :can_add_products, can_update_products = :can_update_products, can_delete_products = :can_delete_products, can_view_products = :can_view_products, can_add_customers = :can_add_customers, can_update_customers = :can_update_customers, can_delete_customers = :can_delete_customers, can_view_customers = :can_view_customers, can_add_subscriptions = :can_add_subscriptions, can_update_subscriptions = :can_update_subscriptions, can_delete_subscriptions = :can_delete_subscriptions, can_view_subscriptions = :can_view_subscriptions, can_add_items = :can_add_items, can_update_items = :can_update_items, can_delete_items = :can_delete_items, can_view_items = :can_view_items, can_add_invoices = :can_add_invoices, can_update_invoices = :can_update_invoices, can_delete_invoices = :can_delete_invoices, can_view_invoices = :can_view_invoices, can_add_payments = :can_add_payments, can_update_payments = :can_update_payments, can_delete_payments = :can_delete_payments, can_view_payments = :can_view_payments, can_view_dashboard = :can_view_dashboard, updated_at = :updated_at
        WHERE id = :id;
    `
	_, err := s.db.NamedExec(query, role)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleStore) DeleteRoleByID(id int) error {
	query := `DELETE FROM roles WHERE id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (s *roleStore) GetRoleByID(id int) (types.Role, error) {
	var role types.Role
	query := `SELECT * FROM roles WHERE id = $1;`
	err := s.db.Get(&role, query, id)
	if err != nil {
		return types.Role{}, err
	}
	return role, nil
}
