package http

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"camarinb2096/cosmetics-shop-go/internal/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

// GetBuyerByID handles the HTTP GET request to find a buyer by ID
func (h *BuyerHandler) GetBuyerByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		IDParam := c.Param("id")
		ID, err := strconv.Atoi(IDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid ID param",
			})
			return
		}
		buyer, err := h.sv.GetBuyerByID(ID)
		if err != nil {
			if errors.Is(err, services.ErrBuyerNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": buyer,
		})
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

// UpdateBuyer handles the HTTP PUT request to update a buyer
func (h *BuyerHandler) UpdateBuyer() gin.HandlerFunc {
	return func(c *gin.Context) {
		IDParam := c.Param("id")
		ID, err := strconv.Atoi(IDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid buyer ID param",
			})
			return
		}

		var buyer entities.BuyerAttributes
		if err := c.ShouldBindJSON(&buyer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid JSON body",
				"error":   err.Error(),
			})
			return
		}

		updatedBuyer, err := h.sv.UpdateBuyer(ID, buyer)
		if err != nil {
			if errors.Is(err, services.ErrBuyerNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "buyer not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update buyer",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": updatedBuyer,
		})
	}
}

// DeleteBuyer handles HTTP DELETE method for delete a buyer
func (h *BuyerHandler) DeleteBuyer() gin.HandlerFunc {
	return func(c *gin.Context) {
		IDParam := c.Param("id")
		ID, err := strconv.Atoi(IDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid ID param",
			})
		}

		err = h.sv.DeleteBuyer(ID)
		if err != nil {
			if errors.Is(err, services.ErrBuyerNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "buyer deleted succesfully",
		})
	}
}
