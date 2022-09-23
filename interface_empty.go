package main

import "C"
import "fmt"

// interface{}
// Peut contenir tous les types

type Person2 struct {
	Name string
	Age  int
}

type Cooker interface {
	Cook()
}

type Chef struct {
	Person2
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("I'm cooking with %d stars\n", c.Stars)
}

func processPerson(i interface{}) {
	switch p := i.(type) {
	case Person2:
		fmt.Printf("We have a Person=%v\n", p)
	case Chef:
		fmt.Printf("We have a Chef=%v, let him cook...\n", p)
		p.Cook()
	default:
		fmt.Printf("Unknown type %T, value=%v\n", p, p)
	}
}

func main() {
	var x interface{} = Person2{"Bob", 10}
	fmt.Printf("x type=%T, data=%v\n", x, x)

	s, ok := x.(string) // Opération de cast en string
	fmt.Printf("Person as string ok? %v. value='%v'\n", ok, s)

	i, ok := x.(int) // Opération de cast en string
	fmt.Printf("Person as int ok? %v. value='%v'\n", ok, i)

	// Processing with Bob
	processPerson(x)

	// Processing with Alice
	x = Chef{
		Stars: 3,
		Person2: Person2{
			Name: "Alice",
			Age:  22,
		},
	}
	processPerson(x)

	processPerson(3)
	processPerson("John")
}
