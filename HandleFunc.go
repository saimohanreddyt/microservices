package main

import (
	"log"
	"net/http"
)

func main() {
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})
	http.ListenAndServe(":9090", nil)

}
