-- Migration Down (Drop Table)
DROP TABLE IF EXISTS coupons;

-- Drop ENUM Types
DROP TYPE IF EXISTS coupon_discount_type;
DROP TYPE IF EXISTS coupon_redemption_type;
DROP TYPE IF EXISTS coupon_linked_to_plan_types;
DROP TYPE IF EXISTS coupon_linked_to_addon_types;