package main

import "fmt"

// Cross compilation
// Compiler le programme pour un OS différent

/*
Maintenir Go sur le serveur
installer les deps sur le serveur
Obligé de copier le code sur le serveur
*/

// Commande
// env GOOS=linux GOARCH=amd64 go build

// Voir la doc ou les sources de GO pour avoir toutes les options possibles

func main() {
	fmt.Println("Hello World")
}
