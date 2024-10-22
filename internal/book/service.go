package service

import (
	"encoding/json"
	"go-crud-app/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	utils.SendErrorResponse(w, http.StatusNotFound, "Book not found")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = params["id"]
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	utils.SendErrorResponse(w, http.StatusNotFound, "Book not found")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	utils.SendErrorResponse(w, http.StatusNotFound, "Book not found")
}
