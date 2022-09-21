package player

type Avatar struct {
	Url string
}

type Player struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}

// New
// Initialisation par fonction
// Va permettre d'initialiser également les champs privés
// ce qui n'est pas faisable depuis un autre package
func New(name string) Player {
	return Player{
		Name:     name,
		password: "defaultPassword",
	}
}
