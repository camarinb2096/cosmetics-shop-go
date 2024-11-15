package services

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"camarinb2096/cosmetics-shop-go/internal/repositories"
	"errors"
)

// BuyerService represents a buyer service
type BuyerService struct {
	rp repositories.BuyerRepositoryInterface
}

// NewBuyerService creates a new instance for buyer service
func NewBuyerService(rp repositories.BuyerRepositoryInterface) BuyerServiceInterface {
	return &BuyerService{
		rp: rp,
	}
}

// GetBuyers retrieves a list of buyers
func (s *BuyerService) GetBuyers() ([]entities.Buyer, error) {
	buyers, err := s.rp.GetBuyers()
	if err != nil {
		return buyers, err
	}

	if len(buyers) == 0 {
		return buyers, errors.New("not buyers found")
	}

	return buyers, nil

}
