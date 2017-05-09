package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/assets", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("templates/coming_soon.html")
		t.Execute(w, nil)
	})
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
