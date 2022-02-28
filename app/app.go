package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-delve/delve/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environemnt variables are not defined")
	}

}

func start() {
	sanityCheck()
	router := mux.NewRouter()

	dbclient := getDbclient()
	customerRepositoryDb := domain.NewcustomerRepositoryDb(dbclient)
	accountRepositoryDb := domain.NewcustomerRepositoryDb(dbclient)

	ch := customerHandlers{service.NewcustomerService(domain.NewcustomerRepositoryDb())}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.GetAllcustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getcustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
func getDbclient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbpasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbport := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbpasswd, dbAddr, dbport, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return client
}
