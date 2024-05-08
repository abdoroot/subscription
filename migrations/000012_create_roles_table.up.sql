-- Create Table
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    company_id INT,
    name VARCHAR(255),
    can_add_products BOOLEAN,
    can_update_products BOOLEAN,
    can_delete_products BOOLEAN,
    can_view_products BOOLEAN, -- New: Permission to view products
    can_add_customers BOOLEAN,
    can_update_customers BOOLEAN,
    can_delete_customers BOOLEAN,
    can_view_customers BOOLEAN, -- New: Permission to view customers
    can_add_subscriptions BOOLEAN,
    can_update_subscriptions BOOLEAN,
    can_delete_subscriptions BOOLEAN,
    can_view_subscriptions BOOLEAN, -- New: Permission to view subscriptions
    can_add_items BOOLEAN,
    can_update_items BOOLEAN,
    can_delete_items BOOLEAN,
    can_view_items BOOLEAN, -- New: Permission to view items
    can_add_invoices BOOLEAN,
    can_update_invoices BOOLEAN,
    can_delete_invoices BOOLEAN,
    can_view_invoices BOOLEAN, -- New: Permission to view invoices
    can_add_payments BOOLEAN,
    can_update_payments BOOLEAN,
    can_delete_payments BOOLEAN,
    can_view_payments BOOLEAN, -- New: Permission to view payments
    can_view_dashboard BOOLEAN, -- New: Permission to view dashboard
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
