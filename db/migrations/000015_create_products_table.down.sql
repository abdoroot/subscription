-- Migration Down (Drop Table)
DROP TABLE IF EXISTS products;

-- Drop ENUM Type
DROP TYPE IF EXISTS product_type_enums;