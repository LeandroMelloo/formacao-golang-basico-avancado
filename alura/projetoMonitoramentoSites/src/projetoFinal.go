package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitaramentos = 3
const delay = 3

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
			imprimeLogs()
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

	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br/"}

	sites := lerSitesDoArquivo()

	for i := 0; i < monitaramentos; i++ {
		fmt.Println("")
		// i => indice do meu slice, site => valor do meu slice
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		// pacote de tempo
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func testaSite(site string) {
	response, err := http.Get(site) // _ = null, ignora o retorno se não quiser utilizar

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo) // lendo o arquivo
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break // saindo do loop for
		}
	}

	arquivo.Close() // fechando o arquivo

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "_ online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
