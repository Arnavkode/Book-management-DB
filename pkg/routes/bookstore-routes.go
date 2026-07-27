package routes

import (
	"github.com/Arnavkode/Book-management-DB/pkg/controllers"
	"github.com/go-chi/chi/v5"
)

var BookstoreRoutes = func(router *chi.Mux) {
	router.Post("/book/", controllers.CreateBook) //create a book
	router.Get("/book/", controllers.GetBook)
	router.Get("/book/{bookId}", controllers.GetBookById)
	router.Put("/book/{bookId}", controllers.UpdateBook)
	router.Delete("/book/{bookId}", controllers.DeleteBook)

}
