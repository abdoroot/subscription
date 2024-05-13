-- Create ENUM Type
CREATE TYPE invoice_items_types_enum AS ENUM ('product', 'subscription');

-- Create Table
CREATE TABLE IF NOT EXISTS invoice_items (
    id SERIAL PRIMARY KEY,
    company_id INT,
    invoice_id INT,
    item_id INT,
    type invoice_items_types_enum,
    qty INT,
    rate DECIMAL(10, 2),
    discount DECIMAL(10, 2),
    amount DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);