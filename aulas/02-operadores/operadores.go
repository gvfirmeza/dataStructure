package main

import "fmt"

func main() {
	var x = 2
	y := 4

	// operadores aritmeticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// unários
	x+=3
	x++
	x--

	// operadores de comparação
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	var a, b = true, false
	// operadores lógicos
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!b)

	// Bitwise
	// * e & 
}