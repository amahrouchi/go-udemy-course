package main

import "fmt"

// Instrumenter
// Suffixe "er" pour les interface courte
type Instrumenter interface {
	Play()
}

type Amplifier interface {
	Connect(amp string)
}

// Pour que les struct soit considérées comme des implémentation d'un interface
// Il suffit d'implémenter les receiver functions définies par l'interface
// Pas besoin de le préciser lors de la déclaration de la struct c'est implicite
// => Duck Typing

type Guitar struct {
	Strings int
}

func (g Guitar) Play() {
	fmt.Printf("Playing guitar with %d strings\n", g.Strings)
}

func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %s", amp)
}

type Piano struct {
	Keys int
}

func (g Piano) Play() {
	fmt.Printf("Playing piano with %d keys\n", g.Keys)
}

func main() {
	var instr Instrumenter
	instr = &Guitar{6}
	instr.Play()
	instr = &Piano{88}
	instr.Play()

	g := Guitar{12}
	g.Play()
	g.Connect("Marshall")

	var g2 Amplifier = Guitar{12}
	g2.Connect("Marshall")
	// Comme le type déclaré est Amplifier, on ne peut pas utiliser la méthode play ici,
	// bien que Guitar implémente en réalité l'interface Instrumenter
	// Ce serait vrai aussi dans le sens inverse si on avait déclaré g2 en Instrumenter
	// on ne pourrait pas utiliser la méthode connect sur cette variable
	//g2.Play()
}
