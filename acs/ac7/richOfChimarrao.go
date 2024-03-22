package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var L, Q float64
	var nomes string

	fmt.Scanln(&N, &L, &Q)
	fmt.Scanln()
	nomes = lerNomes(N)

	participantes := strings.Fields(nomes)
	aguaPorParticipante := Q / float64(len(participantes))
	rico := participantes[len(participantes)-1]

	fmt.Printf("%s %.1f\n", rico, aguaPorParticipante)
}

func lerNomes(N int) string {
	var nomes string
	for i := 0; i < N; i++ {
		var nome string
		fmt.Scan(&nome)
		nomes += nome + " "
	}
	return nomes
}
