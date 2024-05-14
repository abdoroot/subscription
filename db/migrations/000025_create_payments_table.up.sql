-- Create Table
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    company_id INT,
    customer_id INT,
    invoice_id INT,
    amount_received DECIMAL(10, 2),
    payment_date DATE,
    payment_method_id INT,
    reference VARCHAR(255),
    attachments INTEGER[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);