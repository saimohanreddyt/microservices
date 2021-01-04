package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(rw, "Hello %s", d) //Returns the data to the user
	})
	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Good Bye")
	})
	http.ListenAndServe(":9090", nil)

}
