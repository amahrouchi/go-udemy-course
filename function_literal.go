package main

import (
	"fmt"
	"sort"
)

func main() {
	// Definition basique d'une function literal
	func() {
		fmt.Println("Hello")
	}()

	// Avec argument
	func(msg string) {
		fmt.Println(msg)
	}("Hello with args")

	// Utilisation pour trier une liste
	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j]) // Tri par longueur de string
	})
	fmt.Println(people)

	// Le lambda peut etre stocké dans une var
	// La fonction utilise une closure, elle accède au scope de la fonction parente
	// c'est de cette manière qu'elle utilise la var "people"
	people2 := []string{"Angelo", "Meddy", "Jean-Marc"}
	people3 := []string{"Ange", "Meddu", "Jean-Morc"}
	less := func(i, j int) bool {
		return len(people3[i]) < len(people3[j])
	}
	sort.Slice(people2, less)
	fmt.Println(people2)
}
