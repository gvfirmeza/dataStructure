package main

import (
	"fmt"
)

func main() {
	var distancia int
	fmt.Scan(&distancia)

	tempo := (distancia * 2)

	fmt.Printf("%d minutos\n", tempo)
}