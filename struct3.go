package main

import "fmt"

// Embedded struct
// Heritage vs Composition

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User  // Composition ici: Admin est un User
	Level int
}

func main() {
	user := User{
		Name:  "Bob",
		Email: "bob@golang.org",
	}
	fmt.Printf("User=%v\n", user)

	admin := Admin{
		Level: 2,
		// A l'initialisation du User il faut toujours spécifier la clef user
		User: User{
			Name:  "Alice",
			Email: "alice@golang.org",
		},
	}

	// Si l'initialisation est extérieur à la déclaration, il n'est pas nécessaire de spécifier la clef "User"
	//admin.Name = "Alice"
	//admin.Email = "alice@golang.org"

	fmt.Printf("Admin=%v\n", admin)
	fmt.Printf("Admin name=%s, level=%d\n", admin.Name, admin.Level)

	// On peut utiliser Admin comme si c'était un User
}
