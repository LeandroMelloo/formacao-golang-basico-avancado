/*
	O primeiro exemplo é um simples conversor de medidas. Ele aceita como entrada
	uma lista de valores com sua unidade de medida, e produz uma lista de valores convertidos.
	O conversor trabalha apenas com dois tipos de conversões:
	de graus Celsius para Fahrenheit, e de quilômetros pra milhas.

	pacote "fmt", que contém funções para a formatação de strings,

	pacote "strconv" fornece uma grande variedade de funções para a
	conversão de strings para outros formatos e vice-versa.

	pacote "os" possui uma série de operações que lidam com o sistema
	operacional hospedeiro de forma independente, facilitando a escrita de aplicações multiplataformas.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1) // qualquer valor diferente de 0 indica uma execução anormal
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!",
			unidadeDestino)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d não é um número válido!\n",
				v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem,
			valorDestino, unidadeDestino)
	}

}
