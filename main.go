package main

import (
	"fmt" // Format package
	"strings"
	"tests/hello-world/data"
	"tests/hello-world/input"
)

// Run: go run main.go
// Compile: go build

// Lors d'une déclaration de var en dehors d'une fonction, on ne peut pas utiliser le ":=", ce sera toujours "var"
var pi = 3.14

/*
 * Visibilité des symboles en dehors des packages :
 * - privée: 1ere lettre en MINUSCULE
 * - publique: 1ere lettre en MAJUSCULE
 */

func main() {
	input.Keyboard()
	fmt.Println(strings.ToUpper(
		"Hello Gophers! This is a message from the Golang Course."),
	)

	// Typage long
	var age int = 20
	var weight, height int = 80, 190
	fmt.Println(age)
	fmt.Println(weight, height)

	// Typage fort automatique
	var name = "Angelo"
	lastname := "Mahrouchi" // Le : signale une declaration, on ne peut pas déclarer 2 fois
	fmt.Println(name, lastname)
	lastname = "xbH" // On peut ensuite réaffecter une var
	fmt.Println(lastname)

	fmt.Println(pi)

	// Visibilité
	fmt.Println(data.Name, data.Age)
	//fmt.Println(data.password) // Erreur de compilation ici car data.password est privée

	// condition
	age2 := 15
	if age2 < 10 {
		fmt.Printf("petit %d ans\n", age2)
	} else if age2 < 5 {
		fmt.Printf("tout petit %d ans\n", age2)
	} else {
		fmt.Printf("grand %d ans\n", age2)
	}

	// Opérateurs: &&, ||, !
	if age2 > 10 && age < 15 {
		fmt.Printf("alors on est ado %d ans\n", age2)
	}

	// switch avec valeur
	count := 10
	switch count {
	case 1, 2, 3:
		fmt.Println("pas beaucoup")
	case 5:
		fmt.Println("pas mal")
	case 10:
		fmt.Println("trop")
	default:
		fmt.Println("balek")
	}

	// switch avec condition
	switch {
	case count <= 1 && count <= 3:
		fmt.Println("pas beaucoup")

	case count <= 4 && count <= 6:
		fmt.Println("beaucoup")
	default:
		fmt.Println("balek")
	}

	// Cast
	var percent float64 = 2.7
	fmt.Printf("Percent: %f%%\n", percent)
	fmt.Printf("Int Percent: %d%%\n", int(percent))

	percent += 16 // 16 int upcasted to float automatically
	fmt.Printf("Percent upcast: %d%%\n", int(percent))

	n := 42
	f := float64(n) + 0.42
	fmt.Printf("Int upcast to float works: %f\n", f)

	// NOT WORKING: cannot downcast float to int because of information loss on truncation
	//f = n + 0.42
}
