package repositories

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"

	"gorm.io/gorm"
)

// BuyerRepository represents a buyer SQL repository.
// It provides methods to interact with the buyers stored in the database.
type BuyerRepository struct {
	db *gorm.DB // Database connection.
}

// NewBuyerRepository creates a new instance of BuyerRepository.
// It initializes the repository with the given database connection and returns it as an interface.
func NewBuyerRepository(db *gorm.DB) BuyerRepositoryInterface {
	return &BuyerRepository{
		db: db,
	}
}

// GetBuyers retrieves all buyers from the database.
// Returns a slice of Buyer entities and an error, if any occurs during the query.
func (r *BuyerRepository) GetBuyers() ([]entities.Buyer, error) {
	var buyers []entities.Buyer
	err := r.db.Raw(`SELECT * FROM buyers`).Scan(&buyers).Error
	if err != nil {
		return nil, err
	}
	return buyers, nil
}
