package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/davecheney/i2c"
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
	msg := ""
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		log.Println("body", string(body))
		log.Println("body", body)
		msg = string(body[3:])
	case "GET":
		r.ParseForm()
		msg = r.Form.Get("id")
		log.Println("form:", msg)
	}
	//	if r.Method == "POST" {
	//		body, _ := ioutil.ReadAll(r.Body)
	//		log.Println("body", string(body))
	//		log.Println("body", body)
	//		msg = string(body[3:])
	//	}
	//	if r.Method == "GET" {
	//		r.ParseForm()
	//		msg = r.Form.Get("id")
	//		log.Println("form:", msg)
	//	}
	log.Println("method:", r.Method)
	log.Println("msg", msg)

	switch msg {
	case "128":
		fmt.Fprintf(w, "%s", "relays status")
	case "r1":
		fmt.Fprintf(w, "clicked %s!", msg)
	case "r2":
		fmt.Fprintf(w, "clicked %s!", msg)
	case "r3":
		fmt.Fprintf(w, "clicked %s!", msg)
	case "r4":
		fmt.Fprintf(w, "clicked %s!", msg)

	default:
		fmt.Fprintf(w, "%s", "invalid req")
	}
	//		fmt.Fprintf(w, "clicked %s!", msg)

}

func raspberry() {
	i, err := i2c.New(0x20, 1)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css-raspberry-relay-shield"))))
	http.HandleFunc("/ajax/post.html", post)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
