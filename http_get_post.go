package main

import (
	"fmt"
	"log"
	"net/http"
)

// Simple GET request
func helloBoyz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Hello Boyz")
}

// GET request with params
func search(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")

	fmt.Fprintf(w, "term=%v, page=%v\n", t, p)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "http_get_post_login.html")
		return
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed. err=%v\n", err)
			return
		}

		fmt.Fprintf(w, "Login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "Go" && password == "rocks" {
			fmt.Fprintln(w, "You are now logged.")
		} else {
			fmt.Fprintln(w, "Wrong username or password")
		}

		return
	}
}

func main() {
	http.HandleFunc("/", helloBoyz)
	http.HandleFunc("/search", search)
	http.HandleFunc("/login", login)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}
