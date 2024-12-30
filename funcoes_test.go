package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEhMaiorDeIdade(t *testing.T) {
	tests := []struct {
		idade    int
		esperado bool
	}{
		{17, false},
		{18, true},
		{21, true},
		{15, false},
	}

	for _, test := range tests {
		resultado := ehMaiorDeIdade(test.idade)
		if resultado != test.esperado {
			t.Errorf("Erro para idade %d: esperado %v, obtido %v",
				test.idade, test.esperado, resultado)
		}
	}
}

func TestObterIdade(t *testing.T) {
	// Teste para valor válido
	savedStdin := os.Stdin
	defer func() { os.Stdin = savedStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = fmt.Fprintln(w, "25")
	w.Close()

	idade, err := obterIdade()
	if err != nil {
		t.Errorf("Erro ao obter idade: %v", err)
	}
	if idade != 25 {
		t.Errorf("Idade esperada: 25, obtida: %d", idade)
	}

	// Teste para valor inválido
	savedStdin = os.Stdin
	defer func() { os.Stdin = savedStdin }()

	r, w, _ = os.Pipe()
	os.Stdin = r

	_, _ = fmt.Fprintln(w, "abc")
	w.Close()

	_, err = obterIdade()
	if err == nil {
		t.Errorf("Esperava um erro, mas não ocorreu.")
	}
}

func obterIdadeInvalida() (int, error) {
	return 0, fmt.Errorf("digite um número válido")
}
