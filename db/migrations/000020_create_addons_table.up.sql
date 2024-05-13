-- Create ENUM Types
CREATE TYPE addons_billing_type_enums AS ENUM ('once', 'recurring');
CREATE TYPE addons_billing_cycle_enums AS ENUM ('week', 'month', 'year');
CREATE TYPE addons_linked_to_plans_type_enums AS ENUM ('all_plans', 'specific_plans');

-- Create Table
CREATE TABLE IF NOT EXISTS addons (
    id SERIAL PRIMARY KEY,
    company_id INT,
    name VARCHAR(255),
    description VARCHAR(255),
    display_in_widget BOOLEAN,
    pricing_model_id INT,
    unit_id INT,
    billing_type addons_billing_type_enums,
    billing_cycle addons_billing_cycle_enums,
    price DECIMAL(10, 2),
    linked_to_plans addons_linked_to_plans_type_enums,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);