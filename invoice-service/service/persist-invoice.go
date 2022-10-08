package service

import (
	"database/sql"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/google/uuid"
)

const emptyUuid string = "00000000-0000-0000-0000-000000000000"

// Save the invoice data to the database if new invoice. Else, invoice exits update its values.
// TODO: rollback changes incase of error.
func SaveInvoice(db *sql.DB, invoice *models.Invoice) (*models.Invoice, error) {
	var err error
	if invoice.Id.String() == emptyUuid {
		invoice.Id = uuid.New()
		_, err = db.Exec(`INSERT INTO invoices (invoice_id,invoice_number,customer_name, customer_email, due_by, total_amount) VALUES
				($1, $2, $3, $4, $5, $6)`, invoice.Id, invoice.InvoiceNumber, invoice.CustomerName, invoice.Email, invoice.DueDate, invoice.TotalAmount)
		if err == nil {
			err = saveItems(db, invoice.Products, invoice.Id)
		}
	} else {
		var id uuid.UUID
		getInvoiceIdSql := `SELECT invoice_id FROM invoices WHERE invoice_number = $1`
		row := db.QueryRow(getInvoiceIdSql, invoice.InvoiceNumber)
		err = row.Scan(&id)
		if err == nil {
			// Update invoice data and remove and wipe items for this invoice ID and then call the save Items .
			_, err = db.Exec(`UPDATE invoices SET (invoice_number, customer_name, customer_email, due_by, total_amount) =
				($1, $2, $3, $4, $5)`, invoice.InvoiceNumber, invoice.CustomerName, invoice.Email, invoice.DueDate, invoice.TotalAmount)
			if err == nil {
				_, err = db.Exec("DELETE FROM items WHERE invoice_id=$1", id)
				if err == nil {
					saveItems(db, invoice.Products, invoice.Id)
				}
			}
		}
	}
	return invoice, err
}

// Save the line items present in the invoice to the database.
// TODO: rollback changes incase of error.
func saveItems(db *sql.DB, items []models.Item, invoiceId uuid.UUID) error {
	var err error
	for _, item := range items {
		itemId := uuid.New()
		_, err = db.Exec(`INSERT INTO items (item_id, item_name, description, price, invoice_id)
		VALUES ($1, $2, $3, $4, $5)`, itemId, item.Name, item.Description, item.Price, invoiceId)
	}
	return err
}
