package main

import (
	"ExamenGo/aplication"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", aplication.UserRegisterHandler).Methods("POST")
	router.HandleFunc("/login", aplication.UserLoginHandler).Methods("POST")

	log.Println("Server starting on :8090...")
	log.Fatal(http.ListenAndServe(":8090", router))
}
