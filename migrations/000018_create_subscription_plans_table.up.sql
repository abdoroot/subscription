-- Create ENUM Types
CREATE TYPE subscription_plan_type_enums AS ENUM ('service', 'product');
CREATE TYPE subscription_plan_billing_every_enums AS ENUM ('week', 'month', 'year');
CREATE TYPE subscription_plan_billing_cycle_enums AS ENUM ('auto_renewal', 'specific');


-- Create Table
CREATE TABLE IF NOT EXISTS subscription_plans (
    id SERIAL PRIMARY KEY,
    company_id INT,
    product_id INT,
    name VARCHAR(255),
    code VARCHAR(50),
    unit_id INT,
    type subscription_plan_type_enums,
    is_active BOOLEAN,
    price DECIMAL(10, 2),
    billing_every subscription_plan_billing_every_enums,
    billing_every_count INT,
    billing_cycle subscription_plan_billing_cycle_enums,
    billing_cycle_count INT,
    trial_period INT,
    setup_fee DECIMAL(10, 2),
    description TEXT,
    have_features BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);