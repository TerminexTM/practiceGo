package app

import (
	"UdemyREST/domain"
	"UdemyREST/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

//makes sure env variables are set properly or fails to run
//set variables with "export VarName=varString" in terminal
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
	//create db connection
	dbClient := getDbClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDb(dbClient)
	//accountRepositoryDB := domain.NewAccountRepositoryDb(dbClient)
	//wiring for application -- attach to the route in order to reference the appropriate method in the handler
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDB)}
	//defined routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//From handler we receive post request and is passed to newAccount method
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)
	//starting the server. Returns error if issue starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	fmt.Println("Listening on: " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	//"root:Dubu123@@tcp(localhost:3306)/banking"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
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
