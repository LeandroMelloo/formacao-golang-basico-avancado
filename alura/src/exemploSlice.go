package main

import (
	"fmt"
	"reflect"
)

func main() {
	exibirNomes()
}

func exibirNomes() {
	nomes := []string{"Leandro", "Luciana", "Daniel", "Bernado"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem:", len(nomes), "itens")                   // len = tamanho
	fmt.Println("O meu slice tem a capacidade para:", cap(nomes), "itens") // cap = capacidade

	nomes = append(nomes, "Pedro", "Beatriz")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem:", len(nomes), "itens")                   // len = tamanho
	fmt.Println("O meu slice tem a capacidade para:", cap(nomes), "itens") // cap = capacidade
}
