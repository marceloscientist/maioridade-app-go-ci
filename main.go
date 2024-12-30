package main

import "fmt"

func main() {
	idade, err := obterIdade()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	if ehMaiorDeIdade(idade) {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}
}
