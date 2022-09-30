package main

import (
	"encoding/json"
	"fmt"
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

// PasswordJsonBody JSON input
type PasswordJsonBody struct {
	UserIndex         int    `json:"user_index"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
}

var users = []User{
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

// Encode la liste d'utilisateur en JSON
func usersList(w http.ResponseWriter, r *http.Request) {
	// Encodage JSON
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	// On set le content type et on envoie la réponse
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// Modification du pPWD
// Decodage d'un JSON
func updatePassword(w http.ResponseWriter, r *http.Request) {
	var p PasswordJsonBody
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("Parsed struct=%v\n", p)

	// On check l'id du user
	if p.UserIndex < 0 || p.UserIndex > len(users)-1 {
		msg := fmt.Sprintf("Invalid index got user_index=%v. valid range=[0, %v]\n", p.UserIndex, len(users)-1)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}

	// Check password
	u := &users[p.UserIndex] // On recup le pointeur vers le user pour persister le nouveau mdp dans la liste d'origine
	if u.Password != p.OldPassword {
		http.Error(w, "Old password does not match.", http.StatusUnprocessableEntity)
		return
	}

	// Check password confirmation
	if p.NewPassword != p.NewPasswordRepeat {
		http.Error(w, "New password do not match", http.StatusUnprocessableEntity)
		return
	}

	// Update password
	u.Password = p.NewPassword
	fmt.Fprintln(w, "Password updated")
}

func main() {
	http.HandleFunc("/users", usersList)
	http.HandleFunc("/update_password", updatePassword)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}
