// forms.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("forms.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details
		fmt.Println(details)
		empJSON, err := json.MarshalIndent(details, "", " ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("Marshal function output %s\n", string(empJSON))

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		tmpl1 := template.Must(template.ParseFiles("layout.html"))
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl1.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
