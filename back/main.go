package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"backend/app"
	"backend/controllers"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/v1/auth/signup", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/v1/auth/token", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	err := http.ListenAndServe(":80", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Print(err)
	}

}
