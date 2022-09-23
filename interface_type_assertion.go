package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "Miaou!"
}

func main() {
	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}

	// Vérifier le type réel derrière une interface
	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Printf("Type of animal? %T\n", a)

		//t, ok := a.(Dog)
		//fmt.Printf("Type assertion value=%v, ok=%v\n", t, ok)
		if t, ok := a.(Dog); ok {
			fmt.Printf("We have a Dog. color=%v\n", t.color)
		} else {
			fmt.Println("It is not a Dog!")
		}
	}

	// Récupérer le type réel derrière une interface et adapter facilement son code
	fmt.Println("-------------------")
	for _, a := range animals {
		switch v := a.(type) {
		case Dog:
			fmt.Printf("We have a dog, color=%v\n", v.color)
		case Cat:
			fmt.Printf("We have a cat, jumpHeight=%v\n", v.jumpHeight)
		}
	}
}
