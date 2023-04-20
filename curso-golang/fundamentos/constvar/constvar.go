package main

import (
	"fmt"
	m "math" // m é um alias de "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo complilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2) // := estou declarando e inicializando uma varíavel

	fmt.Println("A área da circunferência é", area)

	// outra forma de utilizar var e const
	var (
		a = 1
		b = 2
	)

	const (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// outra forma de utilizar
	var e, f bool = false, true
	fmt.Println(e, f)

	// sintaxe reduzida em Go
	g, h, i := 2, false, "hello world"
	fmt.Println(g, h, i)
}
