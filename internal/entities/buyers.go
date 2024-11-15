package entities

// Buyer represents a buyer who buy on the shop
type Buyer struct {
	ID int
	BuyerAttributes
}

type BuyerAttributes struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Country   string `json:"country"`
}
