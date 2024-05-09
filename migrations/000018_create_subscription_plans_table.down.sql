-- Migration Down (Drop Table)
DROP TABLE IF EXISTS subscription_plans;

-- Drop ENUM Types
DROP TYPE IF EXISTS subscription_plan_type_enums;
DROP TYPE IF EXISTS subscription_plan_billing_every_enums;
DROP TYPE IF EXISTS subscription_plan_billing_cycle_enums;