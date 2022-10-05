package models

import (
	"time"
)

// Model of an Item being purchased.
type Item struct {
	//Id          uuid.UUID   `json:"id"`          // Unique identifier for the item.
	//Unit        unitOfPrice `json:"unit"`        // Unit of cost.
	Name        string  `json:"name"`        // Name of the item
	Description string  `json:"description"` // Description for the item.
	Price       float64 `json:"price"`       // Cost for the item.
}

/*
// Model of Line item- representing item of same id with count.
type LineItem struct {
	Id    uuid.UUID `json:"id"`    // Unique identifier for the line item.
	Items Item      `json:"item"`  // Item in each line of the invoice.
	Cost  float64   `json:"cost"`  // Cost of the item multiplied by count.
}

/*
// Model for an Order-> made up of multiple line items.
type Order struct {
	Id          uuid.UUID `json:"id"`           // Unique identifier for this order.
	LineItems   []Item    `json:"line_items"`   // List of LineItems.
	TotalAmount float64   `json:"total_amount"` // Total cost of the order.
	DueDate     time.Time `json:"due_date"`     // Due date for the invoice.
}
*/
// Model an Invoice.
type Invoice struct {
	Products      []Item  `json:"products"`      // List of products
	InvoiceNumber int64   `json:"invoiceNumber"` // Invoice number.
	TotalAmount   float64 `json:"totalAmount"`   // Total amount of the invoice.
	//Id            uuid.UUID `json:"id"`             // Unique identifier for the invoice.
	CustomerName string `json:"customerName"` // Customer making the purchase.
	Email        string `json:"email"`        // Email id of the customer making the purchase.
	//Order         Order     `json:"order"`          // Order for which the invoice is being created.
	DueDate time.Time `json:"dueByDate"` // Due date for the invoice.
	//CreatedOn     time.Time `json:"created_on"`     // Date when the invoice was created.
	//UpdatedOn     time.Time `json:"updated_on"`     // Date when the invoice was modified.
}

const (
	USDC string = "USDC" // USD Coin
)

func NewInvoice() *Invoice {
	return &Invoice{}
}
