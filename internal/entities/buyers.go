package entities

type Buyer struct {
	ID int
	BuyeerAttributes
}

type BuyeerAttributes struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Country   string `json:"country"`
}
