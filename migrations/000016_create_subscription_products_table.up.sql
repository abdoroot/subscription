-- Create Table
CREATE TABLE IF NOT EXISTS subscription_products (
    id SERIAL PRIMARY KEY,
    company_id INT,
    name VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);