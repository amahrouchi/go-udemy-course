package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name     string  `json:"name"`
	Password string  `json:"-"` // Le `-` permet de ne pas afficher le champ dans le JSON
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country,omitempty"` // omitempty : pour ne pas afficher le pays si le champ est vice
}

func users(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			Name:     "Bob",
			Password: "secret",
			Email:    "bob@golang.org",
			Address: Address{
				Street:  "15 rue Hade",
				City:    "Paris",
				Country: "France",
			},
		},
		{
			Name:     "Alice",
			Password: "super_secret",
			Email:    "alice@golang.org",
			Address: Address{
				Street:  "15 rue Elle",
				City:    "Paris",
				Country: "",
			},
		},
	}

	// Encodage JSON
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	// On set le content type et on envoie la réponse
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/users", users)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}
