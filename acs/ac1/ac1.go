package main

import "fmt"

func calculaMedia(nums ...float64) float64 {
	total := 0.0
	quantidade := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		quantidade++
	}
	return total / float64(quantidade)
}

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
	fmt.Println(calculaMedia(8,5.4,7))
	fmt.Println(verificaParidade(7))
	fmt.Println(minhaPotencia(4,2))
	fmt.Println(converteCelsiusParaFahrenheit(32))
}