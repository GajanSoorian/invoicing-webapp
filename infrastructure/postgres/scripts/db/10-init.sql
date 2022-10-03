CREATE USER invoice_user WITH PASSWORD '123';
CREATE DATABASE invoice_db OWNER invoice_user;
CREATE TABLE IF NOT EXISTS invoices (
	id UUID PRIMARY KEY,
	customer_name VARCHAR ( 50 ) NOT NULL,
	customer_email VARCHAR ( 50 ) UNIQUE NOT NULL,
	created_on TIMESTAMP,
    updated_on TIMESTAMP 
);