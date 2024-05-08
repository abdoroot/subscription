-- Create Table
CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    country_id INT,
    city_id INT,
    is_vat_registered BOOLEAN,
    tax_registration_number VARCHAR(50),
    currency_id INT,
    vat_registration_number VARCHAR(50),
    billing_type billing_type_enum,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);