package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-delve/delve/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())

	} else {
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.code, appError.message)

		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
