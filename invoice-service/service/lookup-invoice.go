package service

import (
	"database/sql"
	"fmt"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/google/uuid"
)

// Fetch Invoice from database based on invoice number.
func FindInvoice(db *sql.DB, invoiceNumber int64) *models.Invoice {
	invoice := models.NewInvoice()
	sqlStatement := `SELECT invoice_id,invoice_number,customer_name, customer_email, due_by, total_amount FROM invoices WHERE invoice_number =$1`
	row := db.QueryRow(sqlStatement, invoiceNumber)
	err := row.Scan(&invoice.Id, &invoice.InvoiceNumber, &invoice.CustomerName, &invoice.Email, &invoice.DueDate, &invoice.TotalAmount)
	switch err {
	case sql.ErrNoRows:
		return nil
	case nil:
		invoice.Products, _ = findItems(invoice.Id, db)
	default:
		fmt.Println(err)
	}
	return invoice
}

// Fetch line items found in invoice.
func findItems(invoiceId uuid.UUID, db *sql.DB) ([]models.Item, error) {
	var items []models.Item
	sqlStatement := `SELECT item_id, item_name, description, price, invoice_id FROM items WHERE invoice_id = $1`
	rows, err := db.Query(sqlStatement, invoiceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := models.NewItem()
		if err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.InvoiceId); err != nil {
			return items, err
		}
		items = append(items, *item)
	}
	if err = rows.Err(); err != nil {
		return items, err
	}
	return items, nil
}
/*
func GetInvoiceId( invoiceNumber int64, db *sql.DB) (uuid.UUID, bool){
	found := false

}*/