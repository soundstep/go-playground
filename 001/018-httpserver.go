package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting server at localhost:8080...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth/", handler2)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	length := len(r.URL.Path)
	remPartOfURL := ""
	if length > len("/earth/") {
		remPartOfURL = r.URL.Path[len("/earth/"):] // get everything after the /earth/ part of the URL
	}
	fmt.Fprintf(w, "Hello Earth %s\n", remPartOfURL)
}
