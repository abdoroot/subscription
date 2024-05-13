-- Create Table
CREATE TABLE IF NOT EXISTS invoices (
    id SERIAL PRIMARY KEY,
    company_id INT,
    customer_id INT,
    order_number VARCHAR(255),
    invoice_date DATE,
    subject VARCHAR(255),
    terms_and_conditions TEXT,
    attachments INTEGER[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);