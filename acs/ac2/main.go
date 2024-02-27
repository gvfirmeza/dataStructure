package main

import (
	"fmt"
	"math"
	"math/rand"
	"ac2/geometria"
)

type Ponto struct {
	x int
	y int
}

func (ponto Ponto) DistanciaOrigem() {
	valor := (math.Sqrt(float64(ponto.x*ponto.x + ponto.y*ponto.y)))
	fmt.Printf("\n%.2f\n", valor)
}

func geraVetor() {

	var vetor [10]int

	for i := 0; i < 10; i++ {
		vetor[i] = rand.Intn(100)
	}

	for i := 0; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}
}

func inverterString(input string) {

	runas := []rune(input)

	n := len(runas)

	for i := 0; i < n/2; i++ {
		runas[i], runas[n-1-i] = runas[n-1-i], runas[i]
	}

	fmt.Println(string(runas))
}

func main() {

	// AC2.1
	geraVetor()

	// AC2.2
	var input string
	fmt.Println("\nDigite uma string:")
	fmt.Scanln(&input)

	inverterString(input)

	// AC2.3
	ponto := Ponto{x: 3, y: 4}

	ponto.DistanciaOrigem()

	// AC2.4
	var largura, altura float64

    	fmt.Print("\nDigite a largura do retângulo: ")
    	fmt.Scanln(&largura)

    	fmt.Print("\nDigite a altura do retângulo: ")
    	fmt.Scanln(&altura)

    	retangulo := geometria.Retangulo{Largura: largura, Altura: altura}

   	area := retangulo.Area()
    	perimetro := retangulo.Perimetro()

    	fmt.Println("\nA área do retângulo é", area)
    	fmt.Println("\nO perímetro do retângulo é", perimetro)
}
