package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Lê e valida a idade a partir da entrada
func obterIdade() (int, error) {
	fmt.Println("Qual a sua idade?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idadeStr := scanner.Text()

	idade, err := strconv.Atoi(idadeStr)
	if err != nil {
		return 0, fmt.Errorf("digite um número válido")
	}
	return idade, nil
}

// Valida se a idade representa maioridade
func ehMaiorDeIdade(idade int) bool {
	return idade >= 18
}
