package main

import (
	"fmt"
)

func main() {
	var N, M, P int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&M)

		precos := make(map[string]float64)
		for j := 0; j < M; j++ {
			var produto string
			var preco float64
			fmt.Scan(&produto, &preco)
			precos[produto] = preco
		}

		fmt.Scan(&P)

		total := 0.0
		for k := 0; k < P; k++ {
			var produto string
			var quantidade int
			fmt.Scan(&produto, &quantidade)
			total += float64(quantidade) * precos[produto]
		}

		fmt.Printf("R$ %.2f\n", total)
	}
}