package main

import (
	"os"
	"text/template"
)

var tpl = `
======= {{.Name}} =======
Grade: {{.Grade}}
Attendance: {{.Attendance}}
School: {{.School}}
`

func main() {
	// Creation d'un template textuel
	t, _ := template.New("student").Parse(tpl)

	// Creation de la struct anonyme qui va servir à regrouper les données à afficher dans le template
	s := struct {
		Name       string
		Grade      int
		Attendance int
		School     string
	}{
		"Bob",
		4,
		8,
		"A",
	}

	// Execution du template avec les données préparées via la struct anonyme
	t.Execute(os.Stdout, s)
}
