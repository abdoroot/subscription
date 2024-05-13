CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    name json NOT NULL,
    code VARCHAR(100) NOT NULL,
    description json
);
