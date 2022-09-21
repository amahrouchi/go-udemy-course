package main

import "fmt"

func UpdateVal(val string) {
	val = "value"
}

func UpdatePtr(ptr *string) {
	*ptr = "pointer"
}

func main() {
	// Int
	i := 1
	var p *int = &i // Syntaxe complete de creation d'un pointer
	fmt.Printf("i=%d\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Println("-----------------")

	// String
	s := "Bob"
	sPtr := &s  // creation d'un pointer vers l'adresse de s
	s2 := *sPtr // Déréférencement de la valeur sur laquelle pointe le pointer sPtr. Une copie de cette valeur est crée et stockée dans s2
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("-----------------")

	// Update par déférencement
	*sPtr = "Alice" // Modification par déréférencement
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("-----------------")

	// Pointer et fonction
	UpdateVal(s) // Ici l'argument est passé par copie donc la modification qui se passe dans la fonction n'a effet que dans la fonction
	// L'affichage ne change donc pas par rapport au précédent
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("-----------------")

	// Pointer et fonction
	UpdatePtr(&s) // On passe l'adresse de "s"
	//UpdatePtr(sPtr) // Mm effet en passant le pointer
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("-----------------")
}
