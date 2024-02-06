package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	Title     string
	Author    string
	Published int
}

func main() {
	fmt.Println("Server started on http://localhost:8000")

	GetBooks := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		books := map[string][]Book{
			"Books": {
				{Title: "Atomic Habit", Author: "James Clear", Published: 2018},
				{Title: "The Power of Habit", Author: "Charles Duhigg", Published: 2012},
				{Title: "Ultralearning", Author: "Scott H. Young", Published: 2019},
			},
		}

		template.Execute(w, books)
	}

	AddBooks := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		publishedStr := r.PostFormValue("published")
		published, _ := strconv.Atoi(publishedStr)

		template := template.Must(template.ParseFiles("index.html"))
		template.ExecuteTemplate(w, "book-list-element", Book{Title: title, Author: author, Published: published})
	}

	http.HandleFunc("/", GetBooks)
	http.HandleFunc("/add-book/", AddBooks)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
