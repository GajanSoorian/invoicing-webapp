package models

import (
	"time"

	"github.com/google/uuid"
)

// Model an Invoice.
type Invoice struct {
	Products      []Item    `json:"products"`      // List of products
	InvoiceNumber int64     `json:"invoiceNumber"` // Invoice number.
	TotalAmount   float64   `json:"totalAmount"`   // Total amount of the invoice.
	Id            uuid.UUID 					   // Unique identifier for the invoice.
	CustomerName  string    `json:"customerName"`  // Customer making the purchase.
	Email         string    `json:"email"`         // Email id of the customer making the purchase.
	DueDate       time.Time `json:"dueByDate"`     // Due date for the invoice.
}

// Model of an Item being purchased. 
type Item struct {
	Id uuid.UUID 							 // Unique identifier for the item.
	InvoiceId uuid.UUID   					 // Invoice Id for this line item.
	Name        string  `json:"name"`        // Name of the item
	Description string  `json:"description"` // Description for the item.
	Price       float64 `json:"price"`       // Cost for the item.
}

const (
	USDC string = "USDC" // USD Coin
)

func NewInvoice() *Invoice {
	return &Invoice{}
}

func NewItem() *Item{
	return &Item{}
}