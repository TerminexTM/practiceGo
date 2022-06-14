package app

import (
	"UdemyREST/domain"
	"UdemyREST/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = "localhost:8000"

func Start() {
	router := mux.NewRouter()

	//wiring for application -- attach to the route in order to reference the appropriate method in the handler
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	//defined routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	fmt.Println("Listening on: " + addr)
	//starting the server. Returns error if issue starting server
	log.Fatal(http.ListenAndServe(addr, router))

}

// router.HandleFunc("/greet", Greet).Methods(http.MethodGet)
// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet) //the [0-9]+ means it accepts only url values that are numeric and any other type will give a 404 error

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post Request Received")
// }
