-- Create ENUM Type
CREATE TYPE product_type_enums AS ENUM ('product', 'service');

-- Create Table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    company_id INT,
    name VARCHAR(255),
    type product_type_enums,
    unit_id INT,
    price DECIMAL(10, 2),
    description TEXT,
    image_attachment_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);