package repositories

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"errors"
)

var (
	ErrBuyerNotFound = errors.New("buyer not found")
)

// BuyerRepositoryInterface represents an interface por buyer sql repository
type BuyerRepositoryInterface interface {
	// GetBuyers retrieves all buyers from the database.
	// Returns a slice of Buyer entities and an error, if any occurs during the query.
	GetBuyers() ([]entities.Buyer, error)
	// CreateBuyer creates a new buyer on DB
	// Returns the object created and an error, if any occurs during the query
	CreateBuyer(buyer entities.BuyerAttributes) (entities.Buyer, error)
	// UpdateBuyer updates a buyer on DB
	// Returns the updated buyer and an error, if any occurs during the query
	UpdateBuyer(ID int, buyer entities.BuyerAttributes) (entities.Buyer, error)
	// GetBuyerByID retrieves a buyer from the database by its ID
	// Returns the buyer and an error if the query fails or no buyer is found
	GetBuyerByID(ID int) (entities.Buyer, error)
}
