-- Create Table
CREATE TABLE IF NOT EXISTS addons_links (
    id SERIAL PRIMARY KEY,
    addon_id INT,
    plan_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);