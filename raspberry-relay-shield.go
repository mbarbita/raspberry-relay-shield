package main

import (
	//	"fmt"
	//"encoding/json"
	"html/template"
	"log"
	"net/http"
	//"io"
	"fmt"
	//"strings"
	"io/ioutil"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Read in the template.
	t, err := template.ParseFiles("templates-raspberry-relay-shield/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")
	}
	t.Execute(w, nil)
	//	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

//type Message struct {
//	Id string
//	//Id string `json:"id"`
//	//Pwd     string `json:"pwd"`
//}

//type Message struct {
//		Name, Text string
//	}

// Handle AJAX Requests
func post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		log.Println("body", string(body))
		log.Println("body", body)
		//	var msg string
		msg := string(body[3:])
		log.Println("method:", r.Method)
		log.Println("msg", msg)
		fmt.Fprintf(w, "clicked %s!", msg)
		//_ = json.Unmarshal(body, &msg)
		//log.Println(msg.Id)
		//var msg Message
		//		decoder := json.NewDecoder(r.Body)
		//	decoder.Decode(&msg)
		// Access data via struct
		//fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
	}
	if r.Method == "GET" {
		r.ParseForm()
		//log.Println("msg:", msg)
		//log.Println("msg.id:", msg.Id)
		log.Println("method:", r.Method)
		log.Println("form:", r.Form.Get("id"))
		fmt.Fprintf(w, "clicked %s!", r.Form.Get("id"))
	}
	//log.Println("form:", r.Form.Get("id"), r.MultipartForm, r.PostForm)
	//log.Println("postformval:", r.PostFormValue("id"))
	//log.Println(r.Body)
	//log.Println(r.Header)
	//log.Println(r.FormValue("id"))

	//	dec := json.NewDecoder(r.Body)
	//
	//	// read open bracket
	//	t, err := dec.Token()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%T: %v\n", t, t)
	//
	//	var m Message
	//	// while the array contains values
	//	for dec.More() {
	//
	//		// decode an array value (Message)
	//		err := dec.Decode(&m)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		fmt.Printf("%v\n", m.Id)
	//	}
	//
	//	// read closing bracket
	//	t, err = dec.Token()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%T: %v\n", t, t)

}

//const jsonStream = `
//		[
//			{"Name": "Ed", "Text": "Knock knock."},
//			{"Name": "Sam", "Text": "Who's there?"},
//			{"Name": "Ed", "Text": "Go fmt."},
//			{"Name": "Sam", "Text": "Go fmt who?"},
//			{"Name": "Ed", "Text": "Go fmt yourself!"}
//		]
//	`

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css-raspberry-relay-shield"))))
	http.HandleFunc("/ajax/post.html", post)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
