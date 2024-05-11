package types

import "time"

type Role struct {
	ID                     int       `db:"id"`
	CompanyID              int       `db:"company_id"`
	Name                   string    `db:"name"`
	CanAddProducts         bool      `db:"can_add_products"`
	CanUpdateProducts      bool      `db:"can_update_products"`
	CanDeleteProducts      bool      `db:"can_delete_products"`
	CanViewProducts        bool      `db:"can_view_products"`
	CanAddCustomers        bool      `db:"can_add_customers"`
	CanUpdateCustomers     bool      `db:"can_update_customers"`
	CanDeleteCustomers     bool      `db:"can_delete_customers"`
	CanViewCustomers       bool      `db:"can_view_customers"`
	CanAddSubscriptions    bool      `db:"can_add_subscriptions"`
	CanUpdateSubscriptions bool      `db:"can_update_subscriptions"`
	CanDeleteSubscriptions bool      `db:"can_delete_subscriptions"`
	CanViewSubscriptions   bool      `db:"can_view_subscriptions"`
	CanAddItems            bool      `db:"can_add_items"`
	CanUpdateItems         bool      `db:"can_update_items"`
	CanDeleteItems         bool      `db:"can_delete_items"`
	CanViewItems           bool      `db:"can_view_items"`
	CanAddInvoices         bool      `db:"can_add_invoices"`
	CanUpdateInvoices      bool      `db:"can_update_invoices"`
	CanDeleteInvoices      bool      `db:"can_delete_invoices"`
	CanViewInvoices        bool      `db:"can_view_invoices"`
	CanAddPayments         bool      `db:"can_add_payments"`
	CanUpdatePayments      bool      `db:"can_update_payments"`
	CanDeletePayments      bool      `db:"can_delete_payments"`
	CanViewPayments        bool      `db:"can_view_payments"`
	CanViewDashboard       bool      `db:"can_view_dashboard"`
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
}

// Validate can be implemented if needed
func (r Role) Validate() bool {
	return true
}
