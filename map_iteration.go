package main

import "fmt"

func main() {
	m := map[string]int{
		"Bob":      10,
		"Alice":    15,
		"Bobbette": 20,
		"John":     7,
	}

	for name, age := range m {
		fmt.Printf("name=%s, age=%d\n", name, age)
	}

	i := 1
	for name := range m {
		fmt.Printf("name=%s\n", name)
		m[name] = i
		i++
	}
	fmt.Println(m)
}
