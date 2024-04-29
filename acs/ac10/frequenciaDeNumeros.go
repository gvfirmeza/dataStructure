package main

import (
	"fmt"
)

func main() {
	var n [2001]int
	var contador, entrada, qtd int
	fmt.Scanf("%d", &qtd)

	for qtd > 0 {
		fmt.Scanf("%d", &entrada)
		n[entrada]++
		qtd--
	}

	for contador = 0; contador < 2001; contador++ {
		if n[contador] > 0 {
			fmt.Printf("%d aparece %d vez(es)\n", contador, n[contador])
		}
	}
}