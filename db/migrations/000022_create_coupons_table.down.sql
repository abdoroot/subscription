-- Migration Down (Drop Table)
DROP TABLE IF EXISTS coupons;

-- Drop ENUM Types
DROP TYPE IF EXISTS coupon_discount_type_enums;
DROP TYPE IF EXISTS coupon_redemption_type_enums;
DROP TYPE IF EXISTS coupon_linked_to_plan_type_enums;
DROP TYPE IF EXISTS coupon_linked_to_addon_type_enums;