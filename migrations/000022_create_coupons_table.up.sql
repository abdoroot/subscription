-- Create ENUM Types
CREATE TYPE coupon_discount_type_enums AS ENUM ('amount', 'percentage');
CREATE TYPE coupon_redemption_type_enums AS ENUM ('once', 'unlimited', 'specific');
CREATE TYPE coupon_linked_to_plan_type_enums AS ENUM ('all', 'specific');
CREATE TYPE coupon_linked_to_addon_type_enums AS ENUM ('all', 'all_recurring', 'all_onetime', 'specific_addons');

-- Create Table
CREATE TABLE IF NOT EXISTS coupons (
    id SERIAL PRIMARY KEY,
    company_id INT,
    name VARCHAR(255),
    discount_value DECIMAL(10, 2),
    discount_type coupon_discount_type_enums,
    redemption_type coupon_redemption_type_enums,
    max_redemptions INT,
    redemption_count INT,
    linked_to_plans coupon_linked_to_plan_type_enums,
    linked_to_addons coupon_linked_to_addon_type_enums,
    expire_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
