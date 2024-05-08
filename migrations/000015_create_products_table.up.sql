-- Create ENUM Type
CREATE TYPE product_types AS ENUM ('product', 'service');

-- Create Table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    company_id INT,
    name VARCHAR(255),
    type product_types,
    unit_id INT,
    price DECIMAL(10, 2),
    description TEXT,
    image_attachment_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);