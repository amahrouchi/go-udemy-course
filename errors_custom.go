package main

import "fmt"

// Pour créer une erreur custom il suffit de créer un struct qui implémente l'interface native "error" de Golang
// Cette interface demande juste de créer un méthode "Error() string" pour notre struct d'erreur custom

const ERR_CODE_LOCKED = 1
const ERR_CODE_CLOSED = 2

type AppError struct {
	Code int
	Err  error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code=%v, error=%v", e.Code, e.Err)
}

func process(input int) error {
	if input == 0 {
		return &AppError{
			Code: ERR_CODE_LOCKED,
			Err:  fmt.Errorf("resource is locked"),
		}
	}

	if input == 1 {
		return &AppError{
			Code: ERR_CODE_CLOSED,
			Err:  fmt.Errorf("resource is closed"),
		}
	}

	return nil
}

func main() {
	// Affichage générique de l'erreur
	err := process(0)
	if err != nil {
		fmt.Println(err)
	}

	// Affichage générique de l'erreur avec un contenu différent
	err = process(1)
	if err != nil {
		fmt.Println(err)
	}

	// L'erreur récupérer ici est une "error" générique, on a donc pas accès au contenu de notre error custom AppError
	// pour ce faire on va caster notre error générique en AppError et coder le comportement particulier
	// que l'on voudrait appliquer dans ce cas.
	err = process(0)
	if err != nil {
		// Affichage de l'erreur générique
		fmt.Println(err)

		// Cast en AppError pour gérer ce cas particulier
		appError, ok := err.(*AppError)
		if ok {
			switch appError.Code {
			case ERR_CODE_LOCKED:
				fmt.Println("Try again later")
			case ERR_CODE_CLOSED:
				fmt.Println("Resource has to be opened again")
			}
		}
	}
}
