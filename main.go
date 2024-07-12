package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hi this is a awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jocrednow@gmail.com\">jocrednow@gmail.com</a></p>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprintln(w, "Page not found")
// 		// or better ->
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintln(w, "Page not found")
		// or better ->
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var pathHandler Router
	fmt.Println("Starting the server on port :3000...")
	http.ListenAndServe("localhost:3000", pathHandler)
}
