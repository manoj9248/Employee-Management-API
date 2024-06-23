package main

import (
	"Employee_crud_mux/endpoints"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter()
	endpoints.EndPoints(s)
	log.Fatal(http.ListenAndServe(":8000", s))
}
