package main

import "fmt"

// slices sont des vues sur des tableaux
// plusieurs slices peuvent pointés sur le mm tableau
// ou mm sur d'autres slices

func main() {
	nums := make(
		[]int, // Type du tableau
		2,     // taille du slice
		3,     // taille du tableau derrière le slice
	)
	nums[0] = 10
	nums[1] = -2
	fmt.Printf("%v, (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	nums = append(nums, 64)
	fmt.Printf("%v, (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	// len est 3 et cap 3

	nums = append(nums, -8)
	fmt.Printf("%v, (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	// ici on dépasse la capacité du tableau lié au slice, la capacité est donc doublée
	// La taille passe donc à 4 (nb d'elements effectivement dans le slice)
	// et la capacité passe à 6 (2 fois la capacité configurée à l'origine)

	// L'affectation rapide pour un slice est identique à celle d'un tableau mais on ne précise pas la taille
	// Implicitement le compilateur comprend qu'on créé un slice et non un tableau
	// Le tableau sou jacent sera d'une capacité égale à la taille du slice créé
	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v, (len=%d, cap=%d)\n", letters, len(letters), cap(letters))

	letters = append(letters, "!")
	fmt.Printf("%v, (len=%d, cap=%d)\n", letters, len(letters), cap(letters))

	// Subslice
	sub1 := letters[0:2] // ou letters[:2]
	sub2 := letters[2:]
	fmt.Printf("sub1=%v, (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2=%v, (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))

	// Si on modifie un subslice, on modifie également les valeurs dans les slices/tableaux d'origine
	sub2[0] = "UP"
	fmt.Printf("letters=%v, (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("sub1=%v, (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2=%v, (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))

	// Pour éviter on peut créér des copies
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	subCopy[0] = "DOWN"
	fmt.Printf("subCopy=%v, (len=%d, cap=%d)\n", subCopy, len(subCopy), cap(subCopy))
	fmt.Printf("sub1=%v, (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	// Le slice d'origine de la copie n'est pas modifié lorsqu'on modifie la copie
}
