package app

import (
	"log"
	"net/http"
)

func Start() {
	//defined routes
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/customers", GetAllCustomers)

	//starting the server. Returns error if issue starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
