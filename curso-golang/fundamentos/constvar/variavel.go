package main

import (
	"fmt"
)

func main() {
	// atribuição de uma variavel sem declarar var, substituindo pelo :
	nome := "Leandro" // string
	idade := 35       // int
	versao := 1.1     // float64
	decisao := true   // bool

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("A variavel decisao retorna", decisao)
}
