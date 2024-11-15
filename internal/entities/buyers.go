package entities

// Buyer represents a buyer who buy on the shop
type Buyer struct {
	ID int
	BuyerAttributes
}

type BuyerAttributes struct {
	FirstName string `json:"first_name"` //FirstName of the buyer
	LastName  string `json:"last_name"`  // LastName of the buyer
	Address   string `json:"address"`    // Address of the buyer
	City      string `json:"city"`       // City of the buyer
	Country   string `json:"country"`    // Country of the buyer
}
