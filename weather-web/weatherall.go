package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HomePageAllHandle(w http.ResponseWriter, r *http.Request)  {
	//city := r.URL.Query().Get("city")
	vmux := mux.Vars(r)
	city := vmux["city"]
	fmt.Fprintln(w, "Weather,", city + "\nHobart 14C shower rain")	
	fmt.Fprintln(w, "New York 0C broken clouds")
	fmt.Fprintln(w, "Kupang 20C clear sky")
	fmt.Fprintln(w, "Nairobi 16C moderate rain")
	fmt.Fprintln(w, "Bangkok 33C few clouds")
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
	r.HandleFunc("/weather/{city}", HomePageAllHandle).Methods("GET")
	http.ListenAndServe(":3000", r)
}