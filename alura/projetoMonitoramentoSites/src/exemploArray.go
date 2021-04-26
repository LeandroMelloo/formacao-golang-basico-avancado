package main

import "fmt"

func main() {
	var frutas [4]string
	frutas[0] = "Abacaxi"
	frutas[1] = "Laranja"
	frutas[2] = "Morango"

	fmt.Println(frutas[3]) // retorna uma string vazia, pois o array tem está limitação
}
