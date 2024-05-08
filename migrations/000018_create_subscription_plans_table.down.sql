-- Migration Down (Drop Table)
DROP TABLE IF EXISTS subscription_plans;

-- Drop ENUM Types
DROP TYPE IF EXISTS subscription_plan_types;
DROP TYPE IF EXISTS subscription_plan_billing_every;
DROP TYPE IF EXISTS subscription_plan_billing_cycle;