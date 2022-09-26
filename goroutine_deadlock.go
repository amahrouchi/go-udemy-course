package main

import (
	"fmt"
	"time"
)

// Deadlock:
// Arrive quand une goroutine essaie d'intéragir avec un channel (R ou W) mais que personne
// n'est de l'autre côté pour échanger.

func hello(c chan string) {
	c <- "Hello there"
	fmt.Println("Hello finished")
}

func hello2(c chan string) {
	c <- "Hello there"
	fmt.Println("Hello finished")
	fmt.Println(<-c)
}

func main() {
	c := make(chan string)

	//go hello(c)              // Essaye d'écrire dans c mais personne n'écoute
	//c <- "message from main" // Essaye d'écrire dans c mais personne n'écoute
	// les 2 goroutines hello() et main() sont bloquées

	// Ici tout se passe car
	// hello2 tente d'écrire dans c
	// main tente de lire dans c avec succès puisque hello2 y écrit
	// ce qui débloque hello2 qui va faire son println puis tenter de lire dans c
	// ce qui va fonctionne puisque main() y écrit avec `c <- "message from main"`
	// ce qui débloque les 2 goroutines main qui tentait d'écrire et hello2 de lire
	go hello2(c)
	fmt.Println(<-c)
	c <- "message from main"

	time.Sleep(time.Second)
}
