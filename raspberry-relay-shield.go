package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	//"github.com/davecheney/i2c"
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
		//log.Println("body", body)
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
	case "startup":
		cha <- 0
		fmt.Fprintf(w, "%v", <-chb)
	case "r1":
		cha <- 1
		fmt.Fprintf(w, "%v", <-chb)
	case "r2":
		cha <- 2
		fmt.Fprintf(w, "%v", <-chb)
	case "r3":
		cha <- 3
		fmt.Fprintf(w, "%v", <-chb)
	case "r4":
		cha <- 4
		fmt.Fprintf(w, "%v", <-chb)

	default:
		fmt.Fprintf(w, "%s", "invalid req")
	}
	//		fmt.Fprintf(w, "clicked %s!", msg)

}

func raspberry() {
	//i, err := i2c.New(0x20, 1)
}

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func execPython(arg int) {
	var cmdOut []byte
	var err error
	cmd := "python"
	//args := []string{"/K","echo","relay.py", string(rs)}
	args := []string{"relay.py", strconv.Itoa(arg)}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pyret := string(cmdOut)
	log.Println("Python script response:", pyret)
	log.Println("Successfully exec python", strconv.Itoa(arg))

}

var cha = make(chan int)
var chb = make(chan int)

func main() {

	go func() {
		rs := 0
		execPython(rs)
		for {
			msga := <-cha
			switch msga {
			case 0:
			case 1:
				if hasBit(rs, 0) {
					rs = clearBit(rs, 0)
				} else {
					rs = setBit(rs, 0)
				}
			case 2:
				if hasBit(rs, 1) {
					rs = clearBit(rs, 1)
				} else {
					rs = setBit(rs, 1)
				}
			case 3:
				if hasBit(rs, 2) {
					rs = clearBit(rs, 2)
				} else {
					rs = setBit(rs, 2)
				}
			case 4:
				if hasBit(rs, 3) {
					rs = clearBit(rs, 3)
				} else {
					rs = setBit(rs, 3)
				}
			default:
				log.Fatal()
			}
			// python call
			execPython(rs)
			//cmd := "cmd"
			//			var cmdOut []byte
			//			var err error
			//			cmd:="python"
			//			//args := []string{"/K","echo","relay.py", string(rs)}
			//			args := []string{"relay.py",strconv.Itoa(rs)}
			//			if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
			//				fmt.Fprintln(os.Stderr, err)
			//				os.Exit(1)
			//			}
			//			pyret := string(cmdOut)
			//			log.Println("Python script response:", pyret)
			//			log.Println("Successfully exec python", strconv.Itoa(rs))
			chb <- rs
			log.Printf("bitwise: %b", rs)
		}
	}()
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css-raspberry-relay-shield"))))
	http.HandleFunc("/ajax/post.html", post)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
