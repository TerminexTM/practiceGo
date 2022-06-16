package domain

import (
	"UdemyREST/errs"
	"UdemyREST/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Sqlx is nice because it handles some error handling and communication
type CustomerRepositoryDB struct {
	db *sqlx.DB
}

//found from looking up mysql connection docs
func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	//custom query logic
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.db.Select(&customers, findAllSql)
		//rows, err = d.db.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.db.Select(&customers, findAllSql, status)
		//rows, err = d.db.Query(findAllSql, status)
	}
	if err != nil {
		logger.Error("Error While Querying Customers Table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customers, nil
	//create a customers object that is based on the dto of Customer
	//---Made obsolete by calling sqlx to query and select in one call
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error While Scanning Customer " + err.Error())
	// 	return nil, errs.NewUnexpectedError("unexpected database error")
	// }

	//---Package Sqlx streamlines this code on line 36------
	//loop through data, map it to Customer and apply it to the customers object
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error While Scanning Customer " + err.Error())
	// 		return nil, errs.NewUnexpectedError("unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	//return the customers object

}

//define a method into the repository layer to find by id
func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) { //returning a custom app error from our error file
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	//---Get method of sqlx streamlines this.
	//row := d.db.QueryRow(customerSql, id)
	var c Customer
	err := d.db.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			//better definition if customer is not found at ID
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			//hide the technical error message for when db is down
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil

}

//helper function to create a db instance user:password@localhost:port/DBname
func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{dbClient}
}
