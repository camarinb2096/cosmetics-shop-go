package services

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"errors"
)

var (
	ErrBuyersNotFound = errors.New("buyers not found")
	ErrBuyerNotFound  = errors.New("buyer not found")
)

type BuyerServiceInterface interface {
	// GetBuyers retrieves a list of buyers
	GetBuyers() ([]entities.Buyer, error)
	// GetBuyerByID retrieves a buyer by ID
	GetBuyerByID(ID int) (entities.Buyer, error)
	// CreateBuyer creates a new buyer
	CreateBuyer(buyer entities.BuyerAttributes) (entities.Buyer, error)
	// UpdateBuyer updates a buyer by their ID
	UpdateBuyer(ID int, buyer entities.BuyerAttributes) (entities.Buyer, error)
	// DeleteBuyer deletes a buyer by their ID
	DeleteBuyer(ID int) error
}
