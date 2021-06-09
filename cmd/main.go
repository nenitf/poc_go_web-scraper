package main

import (
	"errors"
	"fmt"

	vivasul "github.com/nenitf/poabus/pkg/poabus/usecase"
)

func getUrl(linhaNumero string) (string, error) {
	switch linhaNumero {
	case "2601":
		return "http://www.vivasulpoa.com.br/linha?id=51", nil
	case "289":
		return "http://www.vivasulpoa.com.br/linha?id=110", nil
	}

	return "", errors.New("Não foi encontrado exercicio")
}

func main() {
	printLinha("289")
	printLinha("2601")
}

func printLinha(linha string) {
	s := vivasul.NovaBuscaPeloVivaSul()
	url, err := getUrl(linha)

	if err != nil {
		errors.New("Linha inexistente?")
	}
	fmt.Println(s.Execute(url))
}
