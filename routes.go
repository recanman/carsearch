// made by recanman
package main

import "github.com/gorilla/mux"

const SEARCH_BY_ID = "/searches/{id}"

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/makes", GetAllMakes).Methods("GET")

	r.HandleFunc("/searches", CreateSearch).Methods("POST")
	r.HandleFunc(SEARCH_BY_ID, GetSearch).Methods("GET")
	r.HandleFunc(SEARCH_BY_ID, UpdateSearch).Methods("PUT")
	r.HandleFunc(SEARCH_BY_ID, DeleteSearch).Methods("DELETE")
	r.HandleFunc("/searches", GetAllSearches).Methods("GET")
	return r
}
