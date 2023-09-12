// package main

// import (
// 	"net/http"
// )

// func main() {
// 	// fmt.Println("Hello")
// 	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello world"))
// 	})

// 	http.ListenAndServe(":5050", nil)
// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Create a struct to represent your JSON data
type Response struct {
	Message string `json:"message"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define a route for the root path ("/")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a Response struct
		fmt.Println("Hello")
		response := Response{Message: "Hello, World!"}

		// Marshal the struct into JSON format
		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON data as the response
		w.Write(jsonData)
	})

	credentials := handlers.AllowCredentials()

	// Start the HTTP server on port 8081
	fmt.Println("Server is listening on port 8081...")
	http.ListenAndServe(":8081", handlers.CORS(credentials)(router))
}
