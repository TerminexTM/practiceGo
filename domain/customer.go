package domain

//this is a defined business object
//currently we have no business complexity
//so we mov to define repository interface
type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

//helps find all customers from server side
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
