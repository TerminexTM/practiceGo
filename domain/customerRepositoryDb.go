package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	db *sql.DB
}

//found from looking up mysql connection docs
func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	//custom query logic
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	//make query with custom logic
	rows, err := d.db.Query(findAllSql) //information for the method is passed from the receiver "d"
	if err != nil {
		log.Println("Error While Querying Customer Table " + err.Error())
		return nil, err
	}
	//create a customers object that is based on the dto of Customer
	customers := make([]Customer, 0)
	//loop through data, map it to Customer and apply it to the customers object
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error While Scanning Customer " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	//return the customers object
	return customers, nil
}

//helper function to create a db instance user:password@localhost:port/DBname
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	db, err := sql.Open("mysql", "root:Dubu123@@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDB{db}
}
