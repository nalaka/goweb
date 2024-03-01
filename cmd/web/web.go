package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nalaka/goweb/internal/books"
)

const PORT int = 8042

func main() {
	http.HandleFunc("/books", booksHandler)
	fmt.Printf("goweb: listening at port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", PORT), nil))
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", books.BOOKS)
}
