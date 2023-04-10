package main

import (
	"fmt"
)

func main() {
	nome := "Leandro"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

	var comando int    // inicializando a variavel comando
	fmt.Scan(&comando) // & -> identifico o endereço ou ponteiro da variavel

	// testando o valor de uma variavel
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Não conheço este comando digite novamente")
	}
}
