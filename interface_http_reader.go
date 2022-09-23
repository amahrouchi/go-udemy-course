package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Lecture depuis un rq http
	//body, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// Lecture depuis un rq http puis ecriture dans un fichier
	f, _ := os.Create("golang.html")
	defer f.Close()
	io.Copy(f, resp.Body)

	/*
		Ce qui est puissant ici c'est que le reader et le writer utilisé peuvent
		etre remplacé par n'importe que autre type qui implémenterait les interfaces Reader et Writer.
		On pourrait lire depuis l'entrée standard et écrire dans un flux qcq par ex.
	*/
}
