package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoLeandro := ContaCorrente{"Leandro", 334, 345678, 334.99}
	// contaDaLuciana := ContaCorrente{"Luciana", 334, 345678, 334.99}

	fmt.Println(contaDoLeandro)
	// fmt.Println(contaDaLuciana)

	// var contaDoPedro *ContaCorrente // * = ponteiros
	// contaDoPedro = new(ContaCorrente)
	// contaDoPedro.titular = "Pedro"
	// contaDoPedro.saldo = 500

	// fmt.Println(*contaDoPedro)
}
