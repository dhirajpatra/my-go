// basic-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// logging is a middleware function that logs the URL path before calling the provided handler function.
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// foo is a handler function for the "/foo" route, it prints "foo" followed by the requested URL path.
func foo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprintf(w, "foo %s\n", path)
}

// bar is a handler function for the "/bar" route, it prints "bar" followed by the requested URL path.
func bar(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprintf(w, "bar %s\n", path)
}

func main() {
	// Registering the foo and bar handlers with logging middleware for their respective routes.
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	// Listening on port 8080 for incoming HTTP requests.
	http.ListenAndServe(":8080", nil)
}
