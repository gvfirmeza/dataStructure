package main

import "fmt"

func verificaParidade(num int) string {
	if num%2 == 0 {
		return "Par"
	}
	return "Ímpar"
}

func minhaPotencia(base, expoente int) int {
	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	return resultado
}

func converteCelsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	fmt.Println(verificaParidade(7))
	fmt.Println(minhaPotencia(5,8))
	fmt.Println(converteCelsiusParaFahrenheit(32))
}
