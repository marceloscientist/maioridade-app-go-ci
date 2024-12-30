package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	if err := carregarEnv(""); err != nil {
		fmt.Println("Erro ao carregar as variáveis de ambiente:", err)
		return
	}
	run(os.Stdin, os.Stdout)
}

func run(input io.Reader, output io.Writer) {
	idade, err := obterIdade(input)
	if err != nil {
		fmt.Fprintln(output, "Erro:", err)
		return
	}

	if ehMaiorDeIdade(idade) {
		fmt.Fprintln(output, "Você é maior de idade.")
	} else {
		fmt.Fprintln(output, "Você é menor de idade.")
	}
}

func obterIdade(input io.Reader) (int, error) {
	fmt.Println("Qual a sua idade?")
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	idadeStr := scanner.Text()

	idade, err := strconv.Atoi(idadeStr)
	if err != nil {
		return 0, fmt.Errorf("digite um número válido")
	}
	return idade, nil
}

func ehMaiorDeIdade(idade int) bool {
	return idade >= 18
}

func carregarEnv(path ...string) error {
	err := godotenv.Load(path...)
	if err == nil {
		fmt.Println("Variáveis carregadas do .env:")
		fmt.Println("SONAR_TOKEN:", os.Getenv("SONAR_TOKEN"))
		fmt.Println("SONAR_HOST_URL:", os.Getenv("SONAR_HOST_URL"))
	}
	return err
}
