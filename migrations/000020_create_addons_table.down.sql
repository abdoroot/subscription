-- Migration Down (Drop Table)
DROP TABLE IF EXISTS addons;

-- Drop ENUM Types
DROP TYPE IF EXISTS addons_billing_types;
DROP TYPE IF EXISTS addons_billing_cycle;
DROP TYPE IF EXISTS addons_linked_to_plans_types;