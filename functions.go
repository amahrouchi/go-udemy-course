package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Functions:")
	fmt.Printf("add: %d\n", add(10, 2))
	fmt.Printf("add: %d\n", add2(8, 7))

	name, length := multReturn("Alice")
	fmt.Println(length, name)
	name2, _ := multReturn("Jean-Marc") // Ignorer un des retours
	fmt.Println(name2)
}

// Fonction simple (on peut omettre le type de retour si la fonction ne retourne rien)
func add(n int, m int) int {
	return n + m
}

// Fonction avec param de retour nommé
func add2(n int, m int) (sum int) {
	sum = n + m
	return // lorsque la var est nommé dans la déclaration de la fonction, on peut l'omettre dans le return
	// ou return sum (classiquement)
}

// Fonction avec retour multiple
func multReturn(name string) (string, int) {
	return strings.ToLower(name), len(name)
}
