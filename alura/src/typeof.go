package main

import (
	"fmt"
	"reflect"
)

func main() {
	// atribuição de uma variavel sem declarar var, substituindo pelo :
	nome := "Leandro" // string
	idade := 35       // int
	versao := 1.1     // float64
	decisao := true   // bool

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
	fmt.Println("O tipo da variavel decisao é", reflect.TypeOf(decisao))
}
