-- Migration Down (Drop Table)
DROP TABLE IF EXISTS coupon_links;

-- Drop ENUM Type
DROP TYPE IF EXISTS coupon_link_type_enums;