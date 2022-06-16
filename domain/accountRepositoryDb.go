package domain

import (
	"UdemyREST/errs"
	"UdemyREST/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

//accept an account and have all values except ID
func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?);"
	//runs a statement to insert information to DB
	result, err := d.db.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error from Database")
	}
	//after a successful run of this query it will find the last id in the database and count up the id incrementally.
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error from Database")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
