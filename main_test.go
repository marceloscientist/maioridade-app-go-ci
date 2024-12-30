package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
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
	tests := []struct {
		input    string
		esperado int
		erro     bool
	}{
		{"25\n", 25, false},
		{"abc\n", 0, true},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		idade, err := obterIdade(reader)
		if test.erro && err == nil {
			t.Errorf("Esperava um erro para entrada %s, mas não ocorreu.", test.input)
		}
		if !test.erro && idade != test.esperado {
			t.Errorf("Idade esperada: %d, obtida: %d", test.esperado, idade)
		}
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		input    string
		esperado string
	}{
		{"25\n", "Você é maior de idade.\n"},
		{"17\n", "Você é menor de idade.\n"},
		{"abc\n", "Erro: digite um número válido\n"},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		var writer bytes.Buffer

		run(reader, &writer)

		resultado := writer.String()
		if resultado != test.esperado {
			t.Errorf("Para entrada %s, esperado %q, obtido %q", test.input, test.esperado, resultado)
		}
	}
}

func TestCarregarEnv1(t *testing.T) {
	// Criar um arquivo .env temporário para o teste
	envContent := "TEST_ENV=12345\n"
	file, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatalf("Erro ao criar arquivo temporário: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(envContent)
	if err != nil {
		t.Fatalf("Erro ao escrever no arquivo temporário: %v", err)
	}
	file.Close()

	// Carregar o arquivo .env
	os.Setenv("GODOTENV_FILE", file.Name())
	err = godotenv.Load(file.Name())
	if err != nil {
		t.Errorf("Erro ao carregar o .env: %v", err)
	}

	// Validar se a variável foi carregada
	valor := os.Getenv("TEST_ENV")
	if valor != "12345" {
		t.Errorf("Esperado TEST_ENV=12345, mas obtido %s", valor)
	}
}

func TestCarregarEnv(t *testing.T) {
	// Criar um arquivo .env temporário para o teste
	envContent := "TEST_ENV=12345\n"
	file, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatalf("Erro ao criar arquivo temporário: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(envContent)
	if err != nil {
		t.Fatalf("Erro ao escrever no arquivo temporário: %v", err)
	}
	file.Close()

	// Carregar usando a função carregarEnv com o caminho do arquivo
	err = carregarEnv(file.Name())
	if err != nil {
		t.Errorf("Erro ao carregar o .env: %v", err)
	}

	// Validar se a variável foi carregada
	valor := os.Getenv("TEST_ENV")
	if valor != "12345" {
		t.Errorf("Esperado TEST_ENV=12345, mas obtido %s", valor)
	}
}
