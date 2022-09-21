package main

import (
	"fmt"
)

func main() {
	// Syntaxe longue
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while
	n := 0
	for n <= 3 {
		fmt.Println(n)
		n++
	}

	// foreach
	data := [3]int{10, 35, 69}
	for i, v := range data {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}

	for _, i2 := range "string" {
		fmt.Printf("letter=%s\n", string(i2))
	}

	// Boucle infinie
	m := 0
	for {
		m++

		if m%2 == 0 {
			continue // comme dab
		}
		fmt.Println("loop", m)

		if m > 10 {
			break // comme dab
		}

	}
}
