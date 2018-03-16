package main

import (
	"fmt"
	"net/http"
	"mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request)  {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func UsersHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Users Page")
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomePageHandle).Methods("GET")
	r.HandleFunc("/", UsersHandle).Methods("GET")
	return r
}

func main()  {
	http.HandleFunc("/", HomePageHandle)
	http.HandleFunc("/", UsersHandle)
	http.ListenAndServe(":3000", nil)
}