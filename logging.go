package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(
		"logs.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Set the output writer
	log.SetOutput(f)

	// Affiche le message prefixé par la date et l'heure
	log.Println("Hello Gophers")
	log.Println("Output in logs")
}
