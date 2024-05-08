-- Create ENUM Type
CREATE TYPE coupon_link_types AS ENUM ('plan', 'addons');

-- Create Table
CREATE TABLE IF NOT EXISTS coupon_links (
    id SERIAL PRIMARY KEY,
    coupon_id INT,
    link_type coupon_link_types,
    link_to_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);