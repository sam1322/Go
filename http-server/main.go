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

// type Response struct {
// 	Message string `json:"message"`
// }

// func hello(w http.ResponseWriter,r *http.Request){
// 	var resp =
// }

// func main() {
// 	http.HandleFunc("/",hello)
// }

// Create a struct to represent your JSON data
type Response struct {
	Message string `json:"message"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Then, at the start of the handleArticles function, add the following line of code:

// func main() {
//  // fmt.Println("Hello")
//  http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
//      // enableCors(&w)
//      // Create a Response struct
//      response := Response{Message: "successfull "}

//      // Marshal the struct into JSON format
//      jsonData, err := json.Marshal(response)
//      if err != nil {
//          http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
//          return
//      }

//      // Set the Content-Type header to application/json
//      w.Header().Set("Content-Type", "application/json")
//      w.Header().Set("Access-Control-Allow-Origin", "*")
//      // Write the JSON data as the response
//      w.Write(jsonData)

//  })
//  // corsHandler := cors.Default().Handler

//  // // Wrap the CORS handler around your HTTP server
//  // http.Handle("/", corsHandler)
//  // Start the HTTP server on port 8080
//  fmt.Println("Server is listening on port 8081...")
//  err := http.ListenAndServe(":8081", nil)
//  if err != nil {
//      fmt.Println("Error:", err)
//  }
// }

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

	// Create a new CORS handler
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"}, // Change this to the specific origins you want to allow
	// 	AllowedMethods: []string{"GET", "OPTIONS"},
	// })

	// // Use the CORS handler with the router
	// handler := c.Handler(router)
	credentials := handlers.AllowCredentials()

	// Start the HTTP server on port 8080
	fmt.Println("Server is listening on port 8081...")
	http.ListenAndServe(":8081", handlers.CORS(credentials)(router))
	// http.ListenAndServe(":8081", nil)
	// err := http.ListenAndServe(":8081", nil)
	// err := http.ListenAndServe(":8081", handler)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
}
