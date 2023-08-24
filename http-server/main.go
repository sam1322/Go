package main

import (
	"net/http"
)

func main() {
	// fmt.Println("Hello")
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":5050", nil)
}
