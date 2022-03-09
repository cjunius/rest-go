package main

import (
	"net/http"
	"rest-go/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(applicationJsonHeaders)
	r.HandleFunc("/{entityType}", controllers.GetEntities).Methods("GET")
	r.HandleFunc("/{entityType}", controllers.CreateEntity).Methods("POST")
	r.HandleFunc("/{entityType}/{id:[0-9]+}", controllers.GetEntity).Methods("GET")
	r.HandleFunc("/{entityType}/{id:[0-9]+}", controllers.UpdateEntity).Methods("PUT")
	r.HandleFunc("/{entityType}/{id:[0-9]+}", controllers.DeleteEntity).Methods("DELETE")
	http.ListenAndServe("0.0.0.0:8080", r)
}

func applicationJsonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
