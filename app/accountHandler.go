package app

import (
	"UdemyREST/dto"
	"UdemyREST/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

//decode incoming request from app.go - pass to service layer to Validate
func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	//bonus:set customer ID to equal the url value
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	//decode to new account request type 2
	var request dto.NewAccountRequest
	//receive request from client (postman in this case) 1
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//if it fails it is a bad request 3
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		//set customer ID to the url value
		request.CustomerId = customerId
		//on success pass request to newAccount 4
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			//if error stops here and shows error 5
			writeResponse(w, appError.Code, appError.Message)
		} else {
			//then creates the 201 status and creates account 6
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	//get variable information from URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	//decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		//build request object
		request.AccountId = accountId
		request.CustomerId = customerId
	}
	//make the transaction
	account, appError := h.service.MakeTransaction(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, account)
	}
}
