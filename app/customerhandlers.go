package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type customerHandlers struct {
	service service.customerService
}

func (ch *customerHandlers) getAllcustomers(w http.ResponseWriter, r *http.Request) {

	//customers := []Customer{
	//	{Name: "bhavana", city: "warangal", zipcode: "506001"},
	//	{Name: "praneeth", city: "hyderabad", zipcode: "50001"},

	if r.Header.Get("content-type") == "application/xml" {
		w.Header().Add("content-type", "Application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("content-type", "Application/json")
		json.NewEncoder(w).Encode(customers)

	}

	//json.NewEncoder(w).Encode(customers)
}
func (ch *customerHandlers) getcustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.Getcustomer(id)
	if err != nil {
		writeResponse(w, err.code, err.Asmessage())

	} else {
		writeResponse(w, http.StatusOK, customer)

	}
}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
