package main

import (
	"fmt"
	"time"
)

// On définit un taille pour le channel
// L'emetteur bloque quand le channel est plein
// Le recepteur bloque quand le channel est vide
// La capacité d'un channel est definitive

// Une routine qui écrit dans un buffered channel sera bloquée si le channel est plein
// en attendant qu'une autre routine vienne y lire une valeur.

// Un routine qui lit dans un buffered channel ser bloquée si le channel est vide
// et sera en attente qu'un autre routine vienne écrire dans le channel

func write(c chan<- string) {
	names := []string{"Bob", "Alice", "Bobette"}

	//names := []string{"Bob", "Alice", "Bobette", "John"}
	// Avec ce cas, il y a 4 valeurs à écrire dans le channel dont la taille est de 3
	// Donc quand le channel sera plein, la goroutine sera bloqué en attendant qu'une autre routine
	// vienne lire dans le channel et donc libéré une place

	for _, n := range names {
		c <- n // On peut continuer à écrire ici malgré le fait que personne ne lise
		fmt.Printf("Wrote %v to channel (len=%d)\n", n, len(c))
		//time.Sleep(time.Second)
	}
	close(c)
}

func main() {
	c := make(chan string, 3)
	fmt.Printf("Channel data: cap=%v, len=%v\n", cap(c), len(c))

	go write(c)
	time.Sleep(2 * time.Second)

	// La boucle continue tant que des valeurs sont disponibles et que le channel n'est pas fermé
	// Si le channel est vide la goroutine sera bloquée en attendant qu'un autre routine vienne y écrire
	for n := range c {
		fmt.Printf("Read value %v (len=%d)\n", n, len(c))
		time.Sleep(time.Second)
	}
}
