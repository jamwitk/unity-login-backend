package main

import (
	"backend/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/api", controllers.CreateNewUser).Methods("POST")
	r.HandleFunc("/userLogin", controllers.CreateToken).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
