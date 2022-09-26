package main

import (
	"fmt"
	"time"
)

func reader(c chan string) {
	fmt.Println("Start read")
	defer fmt.Println("Stop read")

	// Cette boucle continue tant que le channel reste ouvert
	// Du coup le defer ne pourra se faire qu'après cette fermeture
	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	c := make(chan string)
	go reader(c)

	c <- "Bob"
	c <- "Alice"
	close(c) // On ferme le channel
	// c <- "Bobette" // Va créer un panic error, car le channel est fermé, on ne peut donc plus écrire
	// Par convention, celui qui écrit dans le channel aura la responsabilité de le fermer
	// pour éviter ce genre d'erreur

	time.Sleep(5 * time.Second)
}
