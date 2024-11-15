package entities

// Invoice represents a purchase made by a customer.
type Invoice struct {
	ID string
	InvoiceAttributes
}

type InvoiceAttributes struct {
	CustomerID  string   `json:"customer_id"`          // ID of the customer who made the purchase
	BuyerID     string   `json:"buyer_id"`             // ID of the buyer associated with this invoice
	ProductIDs  []string `json:"product_ids" gorm:"-"` // List of product IDs included in the invoice
	TotalAmount float64  `json:"total_amount"`         // Total amount for the invoice
	Date        string   `json:"date"`                 // Date of the invoice
}
