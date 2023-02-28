package main

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	test := []float64{1.0, 2.0, 3.0, 4.0}
	total := soma(test)

	if total != 10.0 {
		t.Fatalf("Soma incorreta")
	}
}

func TestSubtracao(t *testing.T) {
	test := subtracao(3.8, 2.4)

	if test != 1.4 {
		t.Fatalf("Subtração incorreta")
	}
}

func TestMultriplicao(t *testing.T) {
	tests := []struct {
		Valores   []float64
		Resultado float64
	}{
		{Valores: []float64{5.0, 3.0, 2.0}, Resultado: 30.0},
		{Valores: []float64{-3.0, -10.0, -8.0}, Resultado: -240.0},
	}

	for _, test := range tests {
		total := multiplicacao(test.Valores)

		// if total != test.Resultado {
		// 	t.Fatalf("Multiplicação incorreta")
		// }
		assert.Equal(t, test.Resultado, total)
	}
}

func TestDivisao(t *testing.T) {
	test := divisao(7, 2)

	if test != 3.5 {
		t.Fatalf("Divisão incorreta")
	}
}
