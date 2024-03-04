package main

import (
	"fmt"
)

func resolverHanoi(num int, origem, destino, auxiliar string) {
	if num == 1 {
		fmt.Printf("\nMova o disco da " + origem + " para a " + destino)
		return
	}

	resolverHanoi(num-1, origem, auxiliar, destino)
	
	fmt.Printf("\nMova o disco da " + origem + " para a " + destino)
	
	resolverHanoi(num-1, auxiliar, destino, origem)
}

func buscaSoma(lista []int, valorAlvo int) (int, int) {
	n1 := 0
	n2 := len(lista) - 1

	for n1 < n2 {
		soma := lista[n1] + lista[n2]

		if soma == valorAlvo {
			return lista[n1], lista[n2]
		} else if soma < valorAlvo {
			n1++
		} else {
			n2--
		}
	}
	return -1, -1
}

func main() {

	// AC4.1
	resolverHanoi(4, "Torre 1", "Torre 3", "Torre 2")

	fmt.Println("\n\n---------------------------------------\n")

	// AC4.2
	listaNumeros := []int{1, 2, 4, 6, 8, 14, 23}

	fmt.Println(buscaSoma(listaNumeros, 1))
	fmt.Println(buscaSoma(listaNumeros, 3))
	fmt.Println(buscaSoma(listaNumeros, 18))
	fmt.Println(buscaSoma(listaNumeros, 25))

}
