package service

import (
	"database/sql"
	"fmt"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/google/uuid"
)

// Save the invoice data to the database.
// TODO: rollback changes incase of error.
func SaveInvoice(db *sql.DB, invoice *models.Invoice) (*models.Invoice, error) {
	invoice.Id = uuid.New()
	result, err := db.Exec(`INSERT INTO invoices (invoice_id,invoice_number,customer_name, customer_email, due_by, total_amount) VALUES
				($1, $2, $3, $4, $5, $6)`, invoice.Id, invoice.InvoiceNumber, invoice.CustomerName, invoice.Email, invoice.DueDate, invoice.TotalAmount)
	if err == nil {
		row, _ := result.RowsAffected()
		fmt.Println("Affected number of rows : ", row)
		err = saveItems(db, invoice.Products, invoice.Id)
	}

	return invoice, err
}

// Save the line items present in the invoice to the database.
// TODO: rollback changes incase of error.
func saveItems(db *sql.DB, items []models.Item, invoiceId uuid.UUID) error {
	var err error
	for _, item := range items {
		itemId := uuid.New()
		result, err := db.Exec(`INSERT INTO items (item_id, item_name, description, price, invoice_id)
		VALUES ($1, $2, $3, $4, $5)`, itemId, item.Name, item.Description, item.Price, invoiceId)
		fmt.Println("items id is", itemId)
		if err == nil {
			fmt.Println("items id is inserted")
			row, _ := result.RowsAffected()
			fmt.Println("Affected number of rows : ", row)
		}
	}
	return err
}