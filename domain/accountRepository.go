package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) save(Account) (*Account, *errs.Apperror) {
	sqlInsert = "INSERT INTO accounts(customer_id,opening_date,account_type,amount,status) values (?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.customerId, a.openingdate, a.Accounttype, a.Amount, a.status)
	if err != nil {
		logger.Error("Error while creating new account:" + err.Error())
		return nil, errs.NewUnexpectederror("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account:" + err.Error())
		return nil, errs.NewUnexpectederror("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func NewAccountRepositoryDb(dbclient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbclient}
}
