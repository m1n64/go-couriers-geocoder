package main

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/courier/trackGeo", controllers.SaveCoords).Methods("POST")

	srv := &http.Server{
		Addr:    ":9099",
		Handler: r,
	}

	_ = srv.ListenAndServe()
}
