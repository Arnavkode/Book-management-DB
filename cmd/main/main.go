package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Arnavkode/Book-management-DB/pkg/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()
	routes.BookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server listening")
	log.Fatal(http.ListenAndServe(":8080", r))
}
