package main

import (
	"fmt"
)

func main() {
	
	var numP int
	var capacidadeGarrafa, capacidadeCuia float64
	var listaNomes []string

	fmt.Scan(&numP, &capacidadeGarrafa, &capacidadeCuia)

	for i := 0; i < numP; i++ {
		var nome string
		fmt.Scan(&nome)
		listaNomes = append(listaNomes, nome)
	}

	indiceP := 0
	for capacidadeGarrafa - capacidadeCuia >= 0 {
		capacidadeGarrafa -= capacidadeCuia
		indiceP++
	}

	indiceP = indiceP % numP
	fmt.Printf("%s %.1f\n", listaNomes[indiceP], capacidadeGarrafa)
	
}