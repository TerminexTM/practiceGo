package app

import (
	"UdemyREST/domain"
	"UdemyREST/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring for application -- attach to the route in order to reference the appropriate method in the handler
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//defined routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	//starting the server. Returns error if issue starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

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
