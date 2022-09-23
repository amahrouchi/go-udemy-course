package main

import (
	"fmt"
	"time"
)

// Les goroutines sont des "genre" de threads légers qui sont gérés par le moteur de Go
// et non pas l'OS comme des threads classiques.
// Avantages:
// on peut lancé beaucoup plus de Goroutine que de threads
// Elles demarrent plus rapidement que des threads
// Leur empreinte mémoire est plus faible
// Il n'y pas le changement de contexte que l'OS doit faire pour la création
// Inconvénients:
// Une goroutine a moins de droit qu'un thread et est donc plus limité
// Securite plus faible car l'OS est très méfiant vis a vis des threads mais pas des Goroutine puisqu'elles
// sont gérées par le moeteur de Go.

func longOperation(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished! Took %d s\n", duration)
}

func main() {
	fmt.Println("Starting 1st operation")
	go longOperation(1) // exécuter dans une goroutine séparée

	fmt.Println("Starting 1st operation")
	longOperation(2)

	// Si la goroutine n'est pas terminé à la fonction main (qui est un goroutine également), elle est tuée
	// Pour éviter cela dans notre exemple, on sleep le temps necessaire à la fin de notre main()
	time.Sleep(2 * time.Second)
}
