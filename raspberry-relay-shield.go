package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Read in the template.
	t, err := template.ParseFiles("templates-raspberry-relay-shield/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")
	}
	t.Execute(w, nil)
}

// Handle AJAX Requests
func post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		log.Println("body", string(body))
		log.Println("body", body)
		msg := string(body[3:])
		log.Println("method:", r.Method)
		log.Println("msg", msg)
		fmt.Fprintf(w, "clicked %s!", msg)
	}
	if r.Method == "GET" {
		r.ParseForm()
		log.Println("method:", r.Method)
		log.Println("form:", r.Form.Get("id"))
		fmt.Fprintf(w, "clicked %s!", r.Form.Get("id"))
	}
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css-raspberry-relay-shield"))))
	http.HandleFunc("/ajax/post.html", post)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
