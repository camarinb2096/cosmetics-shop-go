package services

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"errors"
)

var (
	ErrBuyersNotFound = errors.New("buyers not found")
)

type BuyerServiceInterface interface {
	// GetBuyers retrieves a list of buyers
	GetBuyers() ([]entities.Buyer, error)
	// CreateBuyer creates a new buyer
	CreateBuyer(buyer entities.BuyerAttributes) (entities.Buyer, error)
}
