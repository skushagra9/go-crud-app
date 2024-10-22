package main

import (
    "log"
    "net/http"

    "go-crud-app/internal/book"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    book.RegisterBookRoutes(router)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
