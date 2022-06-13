package domain

//THIS IS BASICALLY A MOCK ADAPTER SINCE WE DON'T HAVE A DB SETUP
//create an adapter for the interface port
type CustomerRepositoryStub struct {
	customers []Customer
}

//define the method that now links CustomerRepository Interface to the CustomerRepositoryStub struct
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

//functional responsable for making our dummy customers
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1001", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
