package main

import (
	"fmt"
)

func anosParaUltrapassar(PA int, PB int, G1 float64, G2 float64) int {
	anos := 0
	for PA <= PB {
		PA = int(float64(PA) * (1 + G1/100))
		PB = int(float64(PB) * (1 + G2/100))
		anos = anos + 1
		if anos > 100 {
			return -1
		}
	}
	return anos
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var PA, PB int
		var G1, G2 float64
		fmt.Scan(&PA, &PB, &G1, &G2)

		if G1 > G2 {
			resultado := anosParaUltrapassar(PA, PB, G1, G2)
			if resultado == -1 {
				fmt.Println("Mais de 1 seculo.")
			} else {
				fmt.Println(resultado, "anos.")
			}
		} else {
			fmt.Println("Mais de 1 seculo.")
		}
	}
}