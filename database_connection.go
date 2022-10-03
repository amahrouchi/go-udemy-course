package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // _: forcer l'import à rester explicitement, mm si on ne l'appelle pas (parce qu'on a besoin du driver sqlite pour se connecter)
	"log"
)

func main() {
	// Connexion à la DB
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Ping vers la DB
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Ping DB successful")
}
