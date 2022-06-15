package app

import (
	"UdemyREST/domain"
	"UdemyREST/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//makes sure env variables are set properly or fails to run
func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASS") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment Variable Not Defined")
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	//wiring for application -- attach to the route in order to reference the appropriate method in the handler
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	//defined routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//starting the server. Returns error if issue starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	fmt.Println("Listening on: " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

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
