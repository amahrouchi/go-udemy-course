package benchmark

import (
	"bufio"
	"os"
	"testing"
)

const FILE_PATH = "test.txt"
const LINES_COUNT = 100000
const TEXT = "hello\n"

func BenchmarkWriteFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}

		for i := 0; i < LINES_COUNT; i++ {
			f.WriteString(TEXT)
		}

		f.Close()
	}
}

// Écriture via un buffer : 200 fois plus rapide sur ma machine.
// La boucle n'écrit pas directement sur le disque, elle passe par un buffer mémoire
// qui écrira au moment du flush, On passe donc de 100.000 ecritures disque à 1 seule d'où le gain en
// perf (sachant qu'écrire sur le disque est extrêmement coûteux en terme de temps).
// En réalité, l'OS ne fait pas 100.000 écriture, il utilise également un buffer mais pas de manière aussi
// efficace que si le buffer est créé explicitement.
func BenchmarkWriteFileBuffered(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}

		// Utilisation d'un buffer
		w := bufio.NewWriter(f)

		for i := 0; i < LINES_COUNT; i++ {
			w.WriteString(TEXT)
		}

		w.Flush()
		f.Close()
	}
}
