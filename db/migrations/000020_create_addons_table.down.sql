-- Migration Down (Drop Table)
DROP TABLE IF EXISTS addons;

-- Drop ENUM Types
DROP TYPE IF EXISTS addons_billing_type_enums;
DROP TYPE IF EXISTS addons_billing_cycle_enums;
DROP TYPE IF EXISTS addons_linked_to_plans_type_enums;