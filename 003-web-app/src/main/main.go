/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

SERVEMUX
ServeMux = HTTP request router = multiplexor = mux
compares incoming requests against a list of predefined URL paths,
and calls the associated handler for the path whenever a match is found.

HANDLERS
responsible for writing response headers and bodies.
Almost any type ("object") can be a handler, so long as it satisfies the http.Handler interface.
In lay terms, that simply means it must have a ServeHTTP method with the following signature:
ServeHTTP(http.ResponseWriter, *http.Request)

*/

package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// MyHandler handler
type MyHandler struct {
}

func (p *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println("PATH:", path)

	data, err := ioutil.ReadFile(string(path))
	f, err := os.Open(path)

	if err == nil {

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		bufferedReader := bufio.NewReader(f)

		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)

		if contentType == "video/mp4" {
			bufferedReader.WriteTo(w)
		} else {
			w.Write(data)
		}

	} else {
		log.Println("ERROR:", err)
		w.WriteHeader(404)
		w.Write([]byte("404 MyFriend - " + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	log.Println("Server listening: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/*
We were able to do this because ServeMux also has a ServeHTTP method,
meaning that it too satisfies the Handler interface.

For me it simplifies things to think of a ServeMux as just being a special kind of handler,
which instead of providing a response itself passes the request on to a second handler.

This isn't as much of a leap as it first sounds –
chaining handlers together is fairly commonplace in Go.

from:
http://www.alexedwards.net/blog/a-recap-of-request-handling
*/
