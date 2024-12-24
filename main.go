package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Qual a sua idade?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idadeStr := scanner.Text()

	idade, err := strconv.Atoi(idadeStr)
	if err != nil {
		fmt.Println("Erro: Digite um número válido.")
		return
	}

	if ehMaiorDeIdade(idade) {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade.")
	}
}

func ehMaiorDeIdade(idade int) bool {
	return idade >= 18
}
