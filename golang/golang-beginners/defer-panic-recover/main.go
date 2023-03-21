package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferExample() {
	fmt.Println("deferExample")

	// Getting a response and an optional error
	response, err := http.Get("http://www.google.com/robots.txt")

	// Checking the error if it is not nil so we
	// can log that out and exit the application
	if err != nil {
		log.Fatal(err)
	}
	// Close the body of the response to let the web
	// request know that we're done working with it
	defer response.Body.Close()

	// If the error is not nil, we are going to
	// get the response and take in a stream and
	// parse that out to a string of bytes
	robots, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}

// panicExample
//
// If we open up two terminal and run them,
// we get the application panicking and the reason
// for that is we're trying to access a TCP port
// that's already been blocked. It's been blocked by
// the fact that we're running the application on
// the first terminal.
//
// panic: listen tcp :8080: bind: address already in use
func panicExample() {
	// Registering a function listener on every URL in
	// the application and having a callback that gets called
	// every time a URL comes in that matches the route. This is
	// going to print out the string "Hello Go!".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	// Returning an optional error object
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	deferExample()
	panicExample()
}
