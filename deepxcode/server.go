package main

import "log"
import "time"
import "net/http"
import "github.com/gorilla/mux"
import "deepxcode/handlers"

func main() {
	router := mux.NewRouter()
    handlers.Handle(router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}