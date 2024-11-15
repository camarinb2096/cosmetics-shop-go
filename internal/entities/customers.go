package entities

// Customer represents a customer who shops at the store.
type Customer struct {
	ID string
	CustomerAttributes
}

type CustomerAttributes struct {
	FirstName string `json:"first_name"` // First name of the customer
	LastName  string `json:"last_name"`  // Last name of the customer
	Email     string `json:"email"`      // Email of the customer
	Phone     string `json:"phone"`      // Phone number of the customer
}
