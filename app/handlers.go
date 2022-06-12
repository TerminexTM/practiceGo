package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "11007"},
		{"Rob", "New Delhi", "11005"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
