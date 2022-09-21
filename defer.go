package main

import "fmt"

// defer: très utile pour fermer des ressources en fin de traitement sans risque d'oublier
// fichiers ouverts, connexion bdd, ...

func main() {
	start()
	defer finish()
	// finish() sera exécuté à la fin de la fonction main
	// defer Il permet d'exécuter une instruction à la fin de la fonction

	names := []string{"Bob", "Alice", "Bobette", "John"}
	for _, name := range names {
		fmt.Println("Hello", name)
		defer fmt.Println("Goodbye", name)
		// les defer sont exécuté en LIFO: last in first out
	}

	// le defer finish() est donc exécuté en tout dernier
	// et les defer des names sont executés dans le sens inverses de la déclaration du slice
}

func start() {
	fmt.Println("start")
}

func finish() {
	fmt.Println("finish")
}
