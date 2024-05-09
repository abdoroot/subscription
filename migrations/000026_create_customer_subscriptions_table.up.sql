-- Create Table
CREATE TABLE IF NOT EXISTS customer_subscriptions (
    id SERIAL PRIMARY KEY,
    company_id INT,
    customer_id INT,
    plan_id INT,
    start_date DATE,
    never_expire BOOLEAN,
    expire_after_cycle INT,
    qty INT,
    price DECIMAL(10, 2),
    amount DECIMAL(10, 2),
    reference VARCHAR(255),
    offline_payment BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
