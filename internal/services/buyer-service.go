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
		return buyers, ErrBuyersNotFound
	}

	return buyers, nil

}

// GetBuyerByID retrieves a buyer by ID
func (s *BuyerService) GetBuyerByID(ID int) (entities.Buyer, error) {
	buyer, err := s.rp.GetBuyerByID(ID)
	if err != nil {
		if errors.Is(err, repositories.ErrBuyerNotFound) {
			return entities.Buyer{}, ErrBuyerNotFound
		}
		return entities.Buyer{}, err
	}
	return buyer, nil
}

// CreateBuyer creates a new buyer
func (s *BuyerService) CreateBuyer(buyer entities.BuyerAttributes) (entities.Buyer, error) {
	newBuyer, err := s.rp.CreateBuyer(buyer)
	if err != nil {
		return entities.Buyer{}, err
	}
	return newBuyer, nil
}

// UpdateBuyer updates a buyer by their ID and verifies if the buyer exist
func (s *BuyerService) UpdateBuyer(ID int, buyer entities.BuyerAttributes) (entities.Buyer, error) {
	_, err := s.rp.GetBuyerByID(ID)
	if errors.Is(err, repositories.ErrBuyerNotFound) {
		return entities.Buyer{}, ErrBuyerNotFound
	}

	updatedBuyer, err := s.rp.UpdateBuyer(ID, buyer)
	if err != nil {
		return entities.Buyer{}, err
	}
	return updatedBuyer, nil
}

// DeleteBuyer deletes a buyer by their ID
func (s *BuyerService) DeleteBuyer(ID int) error {
	_, err := s.rp.GetBuyerByID(ID)
	if errors.Is(err, repositories.ErrBuyerNotFound) {
		return ErrBuyerNotFound
	}

	err = s.rp.DeleteBuyer(ID)
	if err != nil {
		return err
	}
	return nil
}
