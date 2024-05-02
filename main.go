package main

import (
	"log"
	"net/http"
)

// home handler function which prints out "Hello World"
// as a response

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Why hello there. Welcome to pagez!"))
}

func pageview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the pageview page!"))
}

func createpage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the pagez create page"))
}

func main() {
	// http mux to initialize new servermux
	// make home function as / handler for url pattern

	router := http.NewServeMux()
	router.HandleFunc("/", home)                   // is a subtree path
	router.HandleFunc("/pagez/view", pageview)     // is a fixed path
	router.HandleFunc("/pagez/create", createpage) // is also a fixed path

	// http.ListenandServe to make a new server
	// 2 parameters. TCP to listen on
	// and router we created. If ListenandServe result in an error 404
	// use log.fatal to error-log and exit
	// errors are ALWAYS none-nil

	log.Println("Starting on :4000")
	error404 := http.ListenAndServe(":4000", router)
	log.Fatal(error404)
}
