package book

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func init() {
	books = append(books, Book{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"})
	books = append(books, Book{ID: "2", Title: "Learning Go", Author: "Jon Bodner"})
}
