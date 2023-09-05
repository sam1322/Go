package main

import (
	"fmt"
	"net/http"
	"os"
)

// request handler function
func hello(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "hello you have requested %v\n", req.URL.Path)
	fmt.Fprint(w, "Welcome to this api")
}

func main() {
	name, dept := "Geeks for Geeks", "CS"
	n, err := fmt.Fprintf(os.Stdout, "%v is a %v portal.\n", name, dept)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "no of bytes were written")

	n1, n2, n3 := 2, 3, 5
	fmt.Fprintf(os.Stdout, "%v + %v = %v\n", n1, n2, n3)

	// registering a request handler to the default HTTP Server
	http.HandleFunc("/", hello)

	// listener function for HTTP Connections
	http.ListenAndServe(":8090", nil)
}
