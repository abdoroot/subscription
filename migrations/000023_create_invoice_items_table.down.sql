-- Migration Down (Drop Table)
DROP TABLE IF EXISTS invoice_items;

-- Drop ENUM Type
DROP TYPE IF EXISTS invoice_items_types_enum;