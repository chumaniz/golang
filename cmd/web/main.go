package main

import (
	"log"
	"net/http"
)

func main() { // function for rendering the renderers
	// http mux to initialize new servermux
	// make home function as / handler for url pattern
	// Do NOT use DeafultServeMux() for production code!
	router := http.NewServeMux()
	router.HandleFunc("/", home)                       // is a subtree path // Method == ANY
	router.HandleFunc("/pagez/pageview", pageview)     // is a fixed path // Method == ANY
	router.HandleFunc("/pagez/pagecreate", createpage) // is also a fixed path //Method == POST

	// http.ListenandServe to make a new server
	// 2 parameters. TCP to listen on
	// and router we created. If ListenandServe result in an error 404
	// use log.fatal to error-log and exit
	// errors are ALWAYS none-nil

	log.Println("Starting on :4000")
	error404 := http.ListenAndServe(":4000", router) // Starts server
	log.Fatal(error404)
}
