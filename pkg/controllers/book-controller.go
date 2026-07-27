package controllers

// CreateBook) //create a book
// 	 controllers.GetBook)
// 	controllers.GetBookById)
// 	controllers.UpdateBook)
// 	 controllers.DeleteBook)

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Arnavkode/Book-management-DB/pkg/models"
	"github.com/Arnavkode/Book-management-DB/pkg/utils"
	"github.com/go-chi/chi/v5"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Printf("Error while parsing", err)
		return
	}
	book := models.GetBookById(int64(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Printf("Error while getting id")
		return
	}

	book := models.DeleteBook(int64(id))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// must have a book struct in my request

	var newBook models.Book

	utils.ParseBody(r, &newBook)

	b := newBook.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// get the id, delete the record, put in new record

	var updatedBook models.Book

	utils.ParseBody(r, &updatedBook)

	queryId := chi.URLParam(r, "bookId")
	id, err := strconv.Atoi(queryId)

	if err != nil {
		fmt.Printf("Error while parsing json")
		return
	}

	bookDetails := models.GetBookById(int64(id))

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	bookDetails.UpdateBook()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(bookDetails)

}
