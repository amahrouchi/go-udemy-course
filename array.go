package main

import "fmt"

func main() {
	// Declaration d'un tableau
	var names [3]string
	fmt.Printf("names=%v (len=%d)\n", names, len(names))

	// Affectation dans un tableau
	names[0] = "Bob"
	names[2] = "Alice"
	fmt.Printf("names=%v (len=%d)\n", names, len(names))
	fmt.Printf("names[2]=%s\n", names[2])

	// Affectation rapide
	odds := [4]int{1, 3, 5, 7}
	fmt.Printf("odss=%v\n (len=%d)", odds, len(odds))

	// Affectation partielle
	odds2 := [4]int{1, 3}
	fmt.Printf("odss=%v\n (len=%d)", odds2, len(odds2))
}
