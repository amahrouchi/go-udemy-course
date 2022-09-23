package main

// Garder les interfaces petites (1 ou 2 fonctions)
// Nommage
// - 1 fonction: suffixe -er
// - éviter de répéter le nom du package dans le nom des interfaces
// - On peut composer les interfaces

/*
 * Ex. d'imbrication
 */

type Saver interface{}
type Loader interface{}
type SaverLoader interface {
	Saver
	Loader
}
