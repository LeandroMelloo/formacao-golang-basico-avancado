package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Println quebra linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	// Convertendo um float64 para imprimir em String
	x := 3.154896
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	// Imprimindo um valor formatado
	fmt.Printf("O valor de x é %.2f", x) // duas casas decimais .2f

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	e := 10
	f := "olá bom dia!"
	/* t -> imprime tipo booleano
	   s -> imprime tipo string
	   d -> imprime tipo decimal
	   f -> imprime tipo float64
	   v -> imprime todos os tipos
	   T -> imprime o tipo da varíavel
	*/
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
	fmt.Println("")
	fmt.Printf("e: %v, %T\n", e, e)
	fmt.Printf("f: %v, %T\n", f, f)

	e = 20 // = (atribuição) é um operador de atribuição e não de declaração que seria := (declaração)
	fmt.Printf("e: %v, %T\n", e, e)
}
