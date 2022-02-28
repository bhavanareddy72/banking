package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type customerRepositoryDb struct {
	client *sql.DB
}

func (d customerRepositoryDb) FindAll() ([]customer, error) {

	FindAllsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
	rows, err := d.client.Query(FindAllsql)
	if err != nil {
		logger.Error("Error while quering custommer table" + err.Error())
		return nil, err
	}


	customers = make([]customer, 0)

	if status==""{
		FindAllsql:="select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err:=d.client.Select(&customers,FindAllsql)
	}else{
		FindAllsql:="select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?"
		err:=d.client.Select(&customers,FindAllsql,status)
		
	}

	if err!=nil{
		logger.Error("Error while quering customer table"+err.Error())
		return nil,errs.NewUnexpectederror("Unexpected database error")
	}
err=sqlx.StructScan(rows,&customers)
if err != nil {
	logger.Error("Error while scanning custommer" + err.Error())
	return nil, errs.NewUnexpectederror("Unexpected database error")

}

	for rows.Next() {
		var c customer
		err := rows.Scan(&c.Id, &c.Name, &c.city, &c.zipcode, &c.DateOfBirth, &c.status)
		if err != nil {
			logger.Error("Error while scanning custommer" + err.Error())
			return nil, err

		}
		customers = append(customers, c)
	}
	return customers, nil
}


func (d customerRepositoryDb)ById(id string)(*customer,*errs.Apperror){
customersql:="select customer_id,name,city,zipcode,Date_of_birth,status from customers where customer_id=?"

var c customer
err:=d.client.Get(&c,customersql,id)

if err!=nil{
	if err==sql.ErrNoRows{
		return nil,errs.NewNotFounderror("customer not found")
	}else{

	logger.Error("Error while scanning customer"+err.Error())
	return nil,errors.NewUnexpectederror("unexpected database error")
}
}
return &c,nil
}

func NewcustomerRepositoryDb(dbclient *sqlx.DB) customerRepositoryDb {
	dbUser:=os.Getenv("DB_USER")
	dbpasswd:=os.Getenv("DB_PASSWD")
	dbAddr:=os.Getenv("DB_ADDR")
	dbport:=os.Getenv("DB_PORT")
	dbName:=os.Getenv("DB_NAME")

	dataSource:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbpasswd,dbAddr,dbport,dbName )
	client, err := sqlx.Open("mysql",dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return customerRepositoryDb{client}

}
}