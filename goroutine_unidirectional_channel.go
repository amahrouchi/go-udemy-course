package main

import (
	"fmt"
	"time"
)

//make(<-chan int) // Read only
//make(chan<- int) // Write only
// Finalement, la syntaxe est la mm que lors de l'utilisation du channel avec sa variable
// -> après : on écrit dans le channel
// -> avant : on lit dans le channel

// longOperation2
// Ici on empèche via la déclaration de la fonction qu'elle puisse lire dans le channel
// Dans le contexte de cette fonction seule les écritures sont autorisées
func longOperation2(done chan<- bool) {
	fmt.Println("Started long operation")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
	done <- true
	//<-done // Cette ligne ne compilera pas
}

func main() {
	done := make(chan bool)
	go longOperation2(done)
	<-done // Ici le channel est utilisé comme un simple drapeau pour attendre la fin de l'execution d'une goroutine
	fmt.Println("Back to main!")
}
