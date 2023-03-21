package main

import (
	"fmt"
	"log"
	"net/http"
)

// Response is what the server sends beck to the user
// Requests is something that the user sends to the server
func formHandler(w http.ResponseWriter, r *http.Request) {
	// User will submit something and there will be a POST
	// request and then that will parse the form
	if err := r.ParseForm(); err != nil {
		// write to w
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")

	// Getting values coming from form.html
	name := r.FormValue("name")
	address := r.FormValue("address")

	// write to w
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Response is what the server sends beck to the user
// Requests is something that the user sends to the server
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Checks if the request is coming from the /hello route or not
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// write to w
	fmt.Fprintf(w, "hello!")
}

func main() {
	// Checking out the static directory
	fileServer := http.FileServer(http.Dir("./static"))
	// Handling root route which is the `/`
	http.Handle("/", fileServer)
	// Handles /form and  will show the form.html
	http.HandleFunc("/form", formHandler)
	// Handles /hello and will show the index.html
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	// This will create the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
