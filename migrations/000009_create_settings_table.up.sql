-- Create Table
CREATE TABLE IF NOT EXISTS settings (
    id SERIAL PRIMARY KEY,
    country_id INT,
    company_id INT,
    name VARCHAR(255),
    value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);