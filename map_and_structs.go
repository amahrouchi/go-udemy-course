package main

import "fmt"

type User2 struct {
	Name string
}

type Key struct {
	Id   int
	Name string
}

func main() {
	// Scalar key
	m := map[string]*User2{
		"HR":  {"Bob"},
		"CEO": {"Alice"},
	}
	fmt.Println(m["HR"])

	hr := m["HR"]
	hr.Name = "John"
	fmt.Println(m["HR"])

	m["CFO"] = &User2{Name: "Bobette"}
	fmt.Println(m["CFO"])

	// Composite keys
	res := make(map[Key]string)
	res[Key{Id: 1, Name: "qwe"}] = "file1"

	keyStruct := Key{Id: 2, Name: "asd"}
	res[keyStruct] = "file2"
	fmt.Println(res)

	// La comparaison est faites entre les valeurs de la struct passé en param et celle effectivement dans la map
	// Si les valeurs coincident alors la clef est supprimée
	delete(res, Key{1, "qwe"})
	fmt.Println(res)
}
