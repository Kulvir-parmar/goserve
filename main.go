package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var users = []string{"JJ"}

func main() {
	fmt.Println("420: 69")

	index := func(w http.ResponseWriter, _ *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		users := map[string][]string{
			"users": users,
		}
		tmpl.Execute(w, users)
	}

	addName := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX baby")
		username := r.FormValue("username")
		users = append(users, username)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "names", username)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/add-name", addName)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
