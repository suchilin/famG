package controllers

import (
	"encoding/json"
	"net/http"
	"backend/app"
	"backend/models"
	u "backend/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(422, "Invalid request"))
		return
	}

	resp := account.Create()

	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(422, "Invalid request"))
		return
	}

	resp, ok := models.Login(account.Username, account.Password)

	if ok {
		app.SetAccessCookie(account.ID, w)
		app.SetRefreshCookie(account.ID, w)
	}
	u.Respond(w, resp)
}
