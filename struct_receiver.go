package main

import "fmt"

type Rect struct {
	Width, Height int
}

// Value receiver: r est passé par copie et non par pointeur/reference
// Consequence: le Rect original sur lequel est appliqué la fonction ne sera pas modifié
func (r Rect) Area() int {
	return r.Width * r.Height
}

// Exemple ici de tentative de modification de value receiver
func (r Rect) doubleSize() {
	r.Width *= 2
	r.Height *= 2

	// Dans la fonction les valeurs sont bien modifiées
	fmt.Println("DoubleSize()", r)
}

// La methode String() sur un struct permet de modifié la manière dont le struct sera affiché dans des print.
// C'est l'équivalent d'un toString() dans d'autres langages.
func (r Rect) String() string {
	return fmt.Sprintf("Rect => w=%d, h=%d", r.Width, r.Height)
}

func main() {
	rect := Rect{2, 4}
	fmt.Printf("Rect area=%d\n", rect.Area())
	fmt.Println(rect)

	rect.doubleSize()
	fmt.Println(rect)
	// On voit ici qu'à l'extérieur de la fonction DoubleSize() les attr n'ont pas été modifiés
	// car rect a été passé par copie à la fonction DoubleSize()
}
