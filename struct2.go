package main

import (
	"fmt"
	"tests/hello-world/player"
)

func main() {
	var p1 player.Player
	p1.Name = "Bob"
	p1.Age = 10

	fmt.Printf("Player 1: %v\n", p1)
	fmt.Printf("p1.Name=%s, p1.Age=%d\n", p1.Name, p1.Age)

	// Initialisation par position
	a := player.Avatar{"http://avatar.jpg"} // Avec cette notation, on doit passer tous les params publics
	fmt.Printf("Avatar: %v\n", a)

	// Initialisation nommée
	p3 := player.Player{
		Name: "Alice",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://avatar3.jpg",
		},
		// Les champs omis prendront les valeurs par défaut standard de Go
	}
	fmt.Printf("Player 3: %v\n", p3)

	// Initialisation par fonction
	p4 := player.New("Bobbette")
	p4.Avatar = a
	fmt.Printf("Player 4: %v\n", p4)
}
