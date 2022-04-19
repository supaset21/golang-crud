package main

import (
	"fmt"
	"log"
	"net/http"

	C "gorilla/controllers"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := 8080
	listenPort := fmt.Sprintf(":%d", port)
	log.Printf("App listen on port %d\n", port)

	router := mux.NewRouter()

	router.HandleFunc("/books", C.ListBookController).Methods("GET")
	router.HandleFunc("/books/{title}", C.ReadBookController).Methods("GET")
	router.HandleFunc("/books/{title}", C.CreateBookController).Methods("POST")
	router.HandleFunc("/books/{title}", C.UpdateBookController).Methods("PUT")
	router.HandleFunc("/books/{title}", C.DeleteBookController).Methods("DELETE")

	http.ListenAndServe(listenPort, router)
}
