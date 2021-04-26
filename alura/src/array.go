package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exiberUsuario()

	for {
		exibeMenu()

		comando := lerComando()

		// testando o valor de uma variavel
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 3:
			fmt.Println("Saindo do Programa...")
			os.Exit(0) // codigo que vai informar pro sistema operacional, que ocorreu tudo bem
		default:
			fmt.Println("Não conheço este comando digite novamente")
			os.Exit(-1) // codigo que vai informar pro sistema operacional, que ocorreu um erro
		}
	}
}

func exiberUsuario() {
	nome := "Leandro"
	versao := 1.2

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {
	var comandoLido int    // inicializando a variavel comando
	fmt.Scan(&comandoLido) // & -> identifico o endereço ou ponteiro da variavel

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br/"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site) // _ = null, ignora o retorno se não quiser utilizar

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
	}
}
