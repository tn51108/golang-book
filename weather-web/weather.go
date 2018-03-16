package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request)  {
	//city := r.URL.Query().Get("city")
	vmux := mux.Vars(r)
	city := vmux["city"]
	if city == "" {
		city = "null"
	}
	fmt.Fprintf(w, "Weather, %s!", city + " shower rain")
}
/*
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", HomePageHandle).Methods("GET")
	return r
}
*/

func main()  {
	
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", HomePageHandle).Methods("GET")
	//r.HandleFunc("/weather/{city}", HomePageHandle).Methods("GET")
	http.ListenAndServe(":3000", r)
}