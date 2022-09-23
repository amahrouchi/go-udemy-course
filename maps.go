package main

import "fmt"

type Vector struct {
	X, Y int
}

func main() {
	// Déclaration
	//var m map[int]bool = make(map[int]bool)
	//var m2 = make(map[int]bool)

	m3 := make(map[int]bool)
	fmt.Println(m3)

	// On peut mettre en clef tout ce qui est comparable:
	// tout sauf slices, func et maps
	// types de bases, struct, ...
	vecs := make(map[Vector]string)
	fmt.Println(vecs)

	m := make(map[string]int)
	fmt.Printf("Map content %v (len=%d)\n", m, len(m))

	m["hello"] = 5
	m["goodbye"] = 10
	fmt.Printf("Map content %v (len=%d)\n", m, len(m))
	fmt.Printf("key=hello, value=%d\n", m["hello"])

	i := m["goodbye"]
	fmt.Println(i)

	// Key exists
	j, ok := m["nope"]
	fmt.Printf("j=%v, ok=%v\n", j, ok)

	m["beatles"] = 2
	if _, ok := m["beatles"]; ok {
		fmt.Println("beatles key exists!")
	}

	// Remove key
	fmt.Printf("Map content %v (len=%d)\n", m, len(m))
	delete(m, "beatles")
	fmt.Printf("Map content %v (len=%d)\n", m, len(m))

	// Direct assignment of maps are done through pointers
	// so the content of both maps will be shared
	m2 := m
	fmt.Printf("m2 content %v (len=%d)\n", m2, len(m2))
	m2["help"] = 44
	fmt.Printf("m2 content %v (len=%d)\n", m2, len(m2))
	fmt.Printf("m content %v (len=%d)\n", m, len(m))

}
