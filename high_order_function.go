package main

import "fmt"

// Fonction qui peuvent etre passé en param
// et être stocké dans des var

type Student struct {
	Name       string
	Grade      int
	Attendance int
}

func filter(students []Student, f func(s Student) bool) []Student {
	var arr []Student
	for _, s := range students {
		if f(s) {
			arr = append(arr, s)
		}
	}
	return arr
}

func genFilterFunc(school string) func(s Student) bool {
	switch school {
	// On accepte tout les étudiants avec la moyenne chez A
	case "A":
		return func(s Student) bool {
			return s.Grade >= 5
		}
	// On accepte les étudiants avec la moyenne et une grande assiduité chez B
	case "B":
		return func(s Student) bool {
			return s.Grade >= 5 && s.Attendance >= 7
		}
	// On accepte tout le monde dans les autres écoles
	default:
		return func(s Student) bool {
			return true
		}
	}

}

func main() {
	s1 := Student{"Bob", 4, 8}
	s2 := Student{"Alice", 8, 4}
	s := []Student{s1, s2}

	// Premier filtre d'étudiant sur le grade
	f := filter(s, func(s Student) bool {
		return s.Grade >= 5
	})
	fmt.Printf("Student filter simple: %v\n", f)

	// Génération d'une fonction en type de retour
	// Tout le monde est pris
	filterFunc := genFilterFunc("Z")
	f = filter(s, filterFunc)
	fmt.Printf("Student filter school Z: %v\n", f)

	// École A (seulement Alice sera prise)
	filterFunc = genFilterFunc("A")
	f = filter(s, filterFunc)
	fmt.Printf("Student filter school A: %v\n", f)

	// École B (personne n'est pris)
	filterFunc = genFilterFunc("B")
	f = filter(s, filterFunc)
	fmt.Printf("Student filter school B: %v\n", f)
}
