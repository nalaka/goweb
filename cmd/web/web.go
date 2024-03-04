package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/nalaka/goweb/internal/books"
)

const PORT int = 8042

//go:embed public
var public embed.FS

//go:embed templates/*.gohtml
var templateFS embed.FS

func main() {
	log.Default()

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(public))))
	mux.Handle("/books/", booksHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "goweb")
	})

	fmt.Printf("goweb: starting server at port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", PORT), mux))
}

type booksHandler struct{}

func (booksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(templateFS, "templates/books.gohtml"))
	tmpl.Execute(w, books.Books)

}
