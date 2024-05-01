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

func main() {
	// http mux to initialize new servermux
	// make home function as / handler for url pattern

	router := http.NewServeMux()
	router.HandleFunc("/", home)

	// http.ListenandServe to make a new server
	// 2 parameters. TCP to listen on
	// and router we created. If ListenandServe result in an error 404
	// use log.fatal to error-log and exit
	// errors are ALWAYS none-nil

	log.Println("Starting on :4000")
	error404 := http.ListenAndServe(":4000", router)
	log.Fatal(error404)
}
