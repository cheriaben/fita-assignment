package main

import (
	d "fita-assignment/database"
	r "fita-assignment/rest"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var err error

func main() {
	db := d.ConnectMySQL()
	log.Info("Connected to database")
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/checkout", r.CheckOut).Methods("POST")
	http.ListenAndServe(":8000", router)
}
