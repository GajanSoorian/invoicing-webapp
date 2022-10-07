CREATE USER invoice_user WITH PASSWORD 'secure_password';
CREATE DATABASE invoice_db OWNER invoice_user;

CREATE TABLE IF NOT EXISTS invoices (
	invoice_id UUID PRIMARY KEY,
	invoice_number BIGINT,
	customer_name VARCHAR ( 50 ),
	customer_email VARCHAR ( 60 ),
	due_by TIMESTAMP,
	total_amount NUMERIC
);

CREATE TABLE IF NOT EXISTS items (
	item_id UUID,
	item_name VARCHAR ( 50 ),
	description VARCHAR ( 200 ),
	price NUMERIC,
	PRIMARY KEY(item_id),
	invoice_id UUID,
	 CONSTRAINT fk_invoice
      FOREIGN KEY(invoice_id) 
	  REFERENCES invoices(invoice_id)
);