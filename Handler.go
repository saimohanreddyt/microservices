package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye")
	})
	http.ListenAndServe(":9090", nil)

}
