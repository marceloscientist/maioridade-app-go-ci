package main

import "testing"

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
