package main

import "fmt"

/*
- Inteiros
	int8, int16, int32, int64
	uint8, uint16, uint32, uint64 → unsigned

	byte → uint8
	rune → int32 → caractere Unicode

- Ponto Flutuante
	float32 ou float64
	float64 → padrão

	complex64 ou complex128

- Texto
	string (cada caractere ocupa 8bits em memória)

- Booleanos
	bool → true or false

- Outros Tipos

ponteiros
*/

var a int
var b = 15
// c := 10 → nao vai funcionar
const PI = 3.1415 // precisa fazer declaracao implicita para constantes

func main() {
	fmt.Println(12.4 + 43.17i)

	msg := `texto em bloco
	top...`

	fmt.Println(msg)

	var x int
	x = 18

	var msg1, msg2 string
	msg1 = "oi"
	msg2 = "tchau"

	var y = 14.5
	z := 20

	fmt.Println(x,y,z,msg1,msg2)

}