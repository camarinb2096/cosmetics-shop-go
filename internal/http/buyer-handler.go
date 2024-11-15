package http

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"camarinb2096/cosmetics-shop-go/internal/services"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BuyerHandler struct {
	sv services.BuyerServiceInterface
}

// NewBuyerHandler initializes a new BuyerHandler with the given service.
func NewBuyerHandler(service services.BuyerServiceInterface) *BuyerHandler {
	return &BuyerHandler{sv: service}
}

// GetBuyers handles the HTTP GET request to retrieve all buyers.
func (h *BuyerHandler) GetBuyers() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyers, err := h.sv.GetBuyers()
		if err != nil {
			if errors.Is(err, services.ErrBuyersNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"buyers": buyers})
	}
}

// CreateBuyer handles the HTTP POST request to create a buyer
func (h *BuyerHandler) CreateBuyer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buyer entities.BuyerAttributes
		err := json.NewDecoder(c.Request.Body).Decode(&buyer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newBuyer, err := h.sv.CreateBuyer(buyer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": newBuyer,
		})
	}
}
