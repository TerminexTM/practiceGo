package domain

//always start here with changes you make to the application
import (
	"UdemyREST/dto"
	"UdemyREST/errs"
)

//this is a defined business object
//currently we have no business complexity
//so we mov to define repository interface
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `json:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

//mapping function for a customer response to map to dto
func (c Customer) ToDto() dto.CustomerResponse {

	//customize dto response string for the status ^
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

//the most important part of sqlx is that it martials response from the database to the domain object -- seen above with the `db`

//helps find all customers from server side
type CustomerRepository interface {
	//status == 1 status == 0 status == "", Active/Inactive/findsall
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
