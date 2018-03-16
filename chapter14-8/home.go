package main

import (
	"time"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/", HomePageHandle).Methods("GET")
	r.HandleFunc("/users", UsersHandle).Methods("GET")
	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		start := time.Now()
		fmt.Printf("Start at %v", start)

		next.ServeHTTP(w,r)

		fmt.Printf("Completed in %v", time.Since(start))
	})
}
func main()  {
	http.HandleFunc("/", HomePageHandle)
	http.HandleFunc("/", UsersHandle)
	http.ListenAndServe(":3000", nil)
}