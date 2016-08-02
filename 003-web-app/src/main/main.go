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
	"fmt"
	"io/ioutil"
	"net/http"
)

//MyHandler
type MyHandler struct {
}

func (p *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path)
	path := r.URL.Path[1:]
	// path := "templates" + r.URL.Path
	data, err := ioutil.ReadFile(string(path))
	if err == nil {
		w.Write(data)
	} else {
		fmt.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("404 MyFriend - " + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/", new(MyHandler))

	fmt.Println("Server listening: http://localhost:8080")
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
