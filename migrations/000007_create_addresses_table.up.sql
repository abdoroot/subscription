-- Create Enum Type
CREATE TYPE address_types AS ENUM ('shipping', 'billing');
-- Create Table
CREATE TABLE IF NOT EXISTS addresses (
    id INT PRIMARY KEY,
    type address_types NOT NULL,
    customer_id INT,   
    country_id INT,
    city_id INT,
    line1 VARCHAR(255),
    line2 VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);