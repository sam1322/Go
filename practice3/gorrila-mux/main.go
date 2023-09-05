package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You have requested the book : %v on page %v \n", title, page)
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// title := vars["title"]
	// page := vars["page"]
	// fmt.Fprintf(w, "You have requested the book : %v on page %v \n", title, page)

	fmt.Fprint(w, "All the books")
}

func main() {

	r := mux.NewRouter()
	// HandleFunc is used instead of HandlerFunc when using custom routers

	// r.HandleFunc("/books/{title}/page/{page}", GetBook).Methods("GET")

	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/all", AllBooks)
	bookrouter.HandleFunc("/{title}/page/{page}", GetBook)

	http.ListenAndServe(":8090", r)

}
