package services

import "camarinb2096/cosmetics-shop-go/internal/entities"

type BuyerServiceInterface interface {
	// GetBuyers retrieves a list of buyers
	GetBuyers() ([]entities.Buyer, error)
}
