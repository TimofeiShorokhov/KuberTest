package controllers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/put", PutData).Methods("GET")
	r.HandleFunc("/get", GetData).Methods("GET")
	r.HandleFunc("/home", Home).Methods("GET")
	r.HandleFunc("/hash", GetHash).Methods("GET")

	return r
}
