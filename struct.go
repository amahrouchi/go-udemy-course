package main

// Struct ne peut contenir que des variables
// Regle de visibilite avec MAJ et MIN pour le struct lui mm et ses variables internes

type Person struct {
	Name      string
	Age       int
	Addr      Address
	birthDate string // Privé : utilisable seulement dans le package courant
}

type Address struct {
	street, city string
}

func main() {
	var P Person
	P.Name = "Angelo"
	P.Age = 36
	P.birthDate = "asdads"
	P.Addr.city = "Le Havre"
}
