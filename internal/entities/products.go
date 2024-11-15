package entities

// Product represents the cosmetic products sold in the store.
type Product struct {
	ID string
	ProductAttributes
}

type ProductAttributes struct {
	Name     string  `json:"name"`     // Name of the product
	Category string  `json:"category"` // Category of the product (e.g., skincare, makeup)
	Price    float64 `json:"price"`    // Price of the product
	Stock    int     `json:"stock"`    // Available stock for the product
}
