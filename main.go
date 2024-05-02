package main

import (
	"log"
	"net/http"
)

// home handler function which prints out "Hello World"
// as a response

func home(writer http.ResponseWriter, request *http.Request) { // function for rendering home page

	if request.URL.Path != "/" {
		http.NotFound(writer, request) // Validating page presence
		return
	}

	writer.Write([]byte("Why hello there. Welcome to pagez!"))

}

func pageview(writer http.ResponseWriter, request *http.Request) { // function for rendering view page

	if request.URL.Path != "/pagez/pageview" {
		http.NotFound(writer, request) // Validating page presence
		return
	}

	writer.Write([]byte("This is the pagez pageview page!"))

}

func createpage(writer http.ResponseWriter, request *http.Request) { // function for rendering create page

	if request.URL.Path != "/pagez/pagecreate" {
		http.NotFound(writer, request)
		return
	}

	if request.Method != "POST" { // validating RESTful method
		writer.Header().Set("Allow", http.MethodPost)
		writer.Header().Set("Content-Type", "application/json")
		http.Error(writer, "Method Not allowed", http.StatusMethodNotAllowed) // No error squiggle. == Best Practice
		return
	}

	writer.Write([]byte("This is the pagez create page"))

}

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
