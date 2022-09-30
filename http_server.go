package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloGophers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Hello Gophers")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Goodbye Gophers")
}

func main() {
	http.HandleFunc("/", helloGophers)
	http.HandleFunc("/goodbye", goodbye)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}
