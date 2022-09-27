package clamp

import (
	"testing"
)

func TestClampValues(t *testing.T) {
	// On prepare le jeu de données via un table (Table Driven Tests)
	// Équivalent à des @dataProvider en PHPUnit
	tests := []struct {
		in       int
		min      int
		max      int
		expected int
	}{
		{-1, 0, 10, 0},
		{0, 0, 10, 0},
		{3, 0, 10, 3},
		{10, 0, 10, 10},
		{11, 0, 10, 10},
	}

	// On test la fonction Clamp en bouclant sur le jeu de données
	for _, tt := range tests {
		v := Clamp(tt.in, tt.min, tt.max)
		if v != tt.expected {
			t.Errorf("Invalid clamp(%v) result. expected=%d got=%d", tt.in, tt.expected, v)
		}
	}
}
