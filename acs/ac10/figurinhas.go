package main

import (
	"fmt"
)

func maxPilha(F1, F2 int) int {
	for F2 != 0 {
		F1, F2 = F2, F1%F2
	}
	return F1
}

func main() {
	var N, F1, F2 int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&F1, &F2)
		tamanhoMaximo := maxPilha(F1, F2)
		fmt.Println(tamanhoMaximo)
	}
}
