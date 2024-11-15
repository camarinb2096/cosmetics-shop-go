package repositories

import "camarinb2096/cosmetics-shop-go/internal/entities"

// BuyerRepositoryInterface represents an interface por buyer sql repository
type BuyerRepositoryInterface interface {
	// GetBuyers retrieves all buyers from the database.
	// Returns a slice of Buyer entities and an error, if any occurs during the query.
	GetBuyers() ([]entities.Buyer, error)
	// CreateBuyer creates a new buyer on DB
	// Returns the object created and an error, if any occurs during the query
	CreateBuyer(buyer entities.BuyerAttributes) (entities.Buyer, error)
}
