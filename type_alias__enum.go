package main

import "fmt"

// Systeme d'ENUM en Golang

type ErrorCode int

const (
	ERR_CODE_OK ErrorCode = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_LOCK
	ERR_CODE_GENERIC
)

// iota: commence à 0 et s'incrémente à chaque appel
// dans un groupe de const seul le 1er iota est nécessaire les autres sont implicites
// car le moteur go comprend ce que l'on souhaite faire ici
// Si on voulait comment à 100 au lieu de 0, il suffirait de faire un "iota + 100" sur la 1ere ligne

// IsCritical
// fonction avec un receiver sur un int
// on voit que grace au type alias on peut definir de nouveau type
// par dessus les type de base et leur associer des fonctions comme on le ferait
// pour des structs ou des objets dans d'autres langage
func (ec ErrorCode) IsCritical() bool {
	return ec == ERR_CODE_LOCK || ec == ERR_CODE_NOT_FOUND
}

// On redéfinit la manière dont le type va s'afficher sous
func (ec ErrorCode) String() string {
	// [...]string: indique que l'on créé un tableau de string
	// dont la taille sera égale au contenu qu'on s'apprête à lui affecter

	return [...]string{
		"ok",
		"not found",
		"locked",
		"generic",
	}[ec]
}

func printErrCode(c ErrorCode) {
	fmt.Printf("code=%v, critical=%v, detail=%v\n", c, c.IsCritical(), c)
}

func main() {
	code := ERR_CODE_OK
	fmt.Println(code)

	printErrCode(code)

	code = ERR_CODE_LOCK
	printErrCode(code)

	code = ERR_CODE_NOT_FOUND
	printErrCode(code)

	// Attention, comme ErrorCode est un int on pourrait écrire
	// printErrCode(10) // Ce code compile
	// Mais risque de planter à l'execution selon la manière dont 10 sera utilisé
	// Ici on plantera car dans la méthode String(), on va aller chercher la clef 10
	// d'un tableau qui n'en contient que 4 (0 à 3)
	// C'est donc une des limitations de ce systeme d'enum en GO qui ne permet pas
	// de valider les valeurs définit dans un type alias
	// Il faudrait potnetiellement créé une fonction qui valide si un int (dans notre cas)
	// est bien dans la liste des valeur disponible d'un ErrorCode
}
