package models

import (
	"time"

	"github.com/google/uuid"
)

// Model of an Item being purchased.
type Item struct {
	Id   uuid.UUID   `json:"id"`   // Unique identifier for the item.
	Cost float64     `json:"cost"` // Cost for the item
	Unit unitOfPrice `json:"unit"` //
}

//Model of Line item- representing item of same id with count .
type LineItem struct {
	Id    uuid.UUID `json:"id"`    // Unique identifier for the line item.
	Items Item      `json:"item"`  // Item in each line of the invoice.
	Count int64     `json:"count"` // Number of items.
	Cost  float64   `json:"cost"`  // Cost of the item multiplied by count.
}

// Model for an Order.
type Order struct {
	Id          uuid.UUID  `json:"id"`           // Unique identifier for this order.
	LineItems   []LineItem `json:"line_items"`   // List of LineItems.
	TotalAmount float64    `json:"total_amount"` // Total cost of the order.
}

// Model an Invoice.
type Invoice struct {
	Id            uuid.UUID `json:"id"`             // Unique identifier for the invoice.
	CustomerName  string    `json:"customer_name"`  // Customer making the purchase.
	CustomerEmail string    `json:"customer_email"` // Email id of the customer making the purchase.
	Order         Order     `json:"order"`          // Order for which the invoice is being created.
	CreatedOn     time.Time `json:"created_on"`     // Date when the invoice was created.
	UpdatedOn     time.Time `json:"updated_on"`     // Date when the invoice was modified.
}

type unitOfPrice string

const (
	USDC string = "USDC" // USD Coin
)

func (lineItem *LineItem) TotalCost() {
	lineItem.Cost = lineItem.Items.Cost * float64(lineItem.Count)
}

func New() *Invoice {
	return &Invoice{}
}

// "Behavior" method of Invoice struct, for adding items to an invoice
func (order *Order) AddLineItem(lineItem LineItem) {
	order.LineItems = append(order.LineItems, lineItem)
}

/*
// Model for Order
type Order struct {
    Id    uuid.UUID `json:"id"`   // Unique identifier for the order.
    Items []Item    `json:"items` // List of items present in the order.
}
*/
