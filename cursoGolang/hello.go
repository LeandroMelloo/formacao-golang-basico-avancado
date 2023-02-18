package main

// pacote
import "fmt"

// func main() -> porta de entrada e porta de saída
// ...interface{} -> posso passar quantos elementos eu quiser, e do tipo que eu quiser
// _ ou blank identifier -> ignora o valor e descarta
func main() {
	_, error := fmt.Println("Hello world!", "Teste", true, 1000)
	fmt.Println(error)
}