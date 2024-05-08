-- Create Enum Type
CREATE TYPE customer_types AS ENUM ('individual', 'company');
-- Create Table
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    type customer_types NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    company_name VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);