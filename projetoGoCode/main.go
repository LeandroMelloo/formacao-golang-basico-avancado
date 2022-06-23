package main

import "fmt"

func main() {
	// Por convenção no Go as variaveis vem no topo da função
	var nome = "Conferência Go"
	var ticketsSobrando = 100
	const ticketsConferencia = 100
	nomeCamelCase := "Conferência Go 1"

	// comentario no Go
	fmt.Println("Olá Mundo")
	fmt.Println("Bem vindo à", nome, "a melhor conferência do mundo sobre Go")
	fmt.Println(nomeCamelCase)
	fmt.Printf("Temos um total de %v tickets e %v tickets disponíveis para compra. Garanta logo sua vaga!", ticketsConferencia, ticketsSobrando)
}
