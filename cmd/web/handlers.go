package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home handler function which prints out "Hello World"
// as a response

func home(writer http.ResponseWriter, request *http.Request) { // function for rendering home page

	if request.URL.Path != "/" {
		http.NotFound(writer, request) // Validating page presence (AKA a Router)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...) // This is to reads html when project runs
	if err != nil {                          // If there is an error validation
		log.Println(err.Error())
		http.Error(writer, "Internal server error", 500) // What will be printed out in place of the html
		return
	}

	// Execute() makes the file path the response body
	err = ts.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal server error", 500) // This executes the page
	}

}

// function for rendering view page(AKA a handler)

func pageview(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/pagez/pageview" {
		http.NotFound(writer, request) // Validating page presence
		return
	}

	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, "This is the pagez pageview with id: %d ", id) // I belive this confirms that this is the specific page
	// You are looking for.

	writer.Write([]byte("This is the pagez pageview page!"))

}

func createpage(writer http.ResponseWriter, request *http.Request) { // function for rendering create page(AKA a handler)

	if request.URL.Path != "/pagez/pagecreate" {
		http.NotFound(writer, request)
		return
	}

	if request.Method != "POST" { // validating RESTful method
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Method Not allowed", http.StatusMethodNotAllowed) // No error squiggle. == Best Practice
		return
	}

	writer.Write([]byte("This is the pagez create page"))

}
