// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em GO são organizados em pacotes e para usá-los é necessário declarar um ou vários imports
*/
import "fmt"

// A porta de entrada de um programa Go é a função main
func main() {

	x := 16
	y := "String"
	z := true

	fmt.Print("Primeiro ")
	fmt.Println("Programa!") // Println -> quebra linha

	numerodebytes, erros := fmt.Println("Hello world!")

	fmt.Println(numerodebytes, erros)
	fmt.Println(x, y, z)

	/*
		Sobre comentários...

		1) Priorize código legível e faça comentários que agrega valor!
		2) Evite comentários óbvios
		3) Durante o curso abuse dos comentários
	*/
}
