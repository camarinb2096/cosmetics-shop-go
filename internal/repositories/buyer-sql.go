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

// CreateBuyer creates a new buyer on DB
// Returns the object created and an error, if any occurs during the query
func (r *BuyerRepository) CreateBuyer(buyer entities.BuyerAttributes) (entities.Buyer, error) {
	newBuyer := entities.Buyer{
		BuyerAttributes: buyer,
	}

	result := r.db.Create(&newBuyer)
	if result.Error != nil {
		return entities.Buyer{}, result.Error
	}

	return newBuyer, nil
}

// UpdateBuyer updates a buyer in the database
// Returns the updated buyer and an error, if any occurs during the query
func (r *BuyerRepository) UpdateBuyer(ID int, buyer entities.BuyerAttributes) (entities.Buyer, error) {
	updatedBuyer := entities.Buyer{
		ID:              ID,
		BuyerAttributes: buyer,
	}

	if err := r.db.Save(&updatedBuyer).Error; err != nil {
		return entities.Buyer{}, err
	}

	return updatedBuyer, nil
}

// GetBuyerByID retrieves a buyer from the database by its ID
// Returns the buyer and an error if the query fails or no buyer is found
func (r *BuyerRepository) GetBuyerByID(ID int) (entities.Buyer, error) {
	var buyer entities.Buyer

	err := r.db.First(&buyer, ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.Buyer{}, ErrBuyerNotFound
		}
		return entities.Buyer{}, err
	}

	return buyer, nil
}

// DeleteBuyer deletes a buyer from the database by their ID
// Returns an error if the query fails
func (r *BuyerRepository) DeleteBuyer(ID int) error {
	result := r.db.Delete(&entities.Buyer{}, ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
