package app

import (
	"UdemyREST/service"
	"encoding/json"
	"net/http"
)

//connect our rest handlers with the service port
type CustomerHandlers struct {
	service service.CustomerService
}

//this is a concrete implementation of the handler

// func Greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "hello world!!")
// }

//reference our struct via a receiver
func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Ashish", "New Delhi", "11007"},
	// 	{"Rob", "New Delhi", "11005"},
	// }
	//instead of hard coding the response body we can reference a response stub now!
	customers, _ := ch.service.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
