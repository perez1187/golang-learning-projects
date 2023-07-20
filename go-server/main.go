package main

import (
	"fmt"
	"log"
	"net/http"
)

// we have 3 route, / , hello, form

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succes")
	// we take values from form
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name - %s\n", name)
	fmt.Fprintf(w, "Address - %s \n", address)
}

// r- request
// * it is a pointer
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// we check correct path
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//we check method
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer) // defaoult route(?)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
