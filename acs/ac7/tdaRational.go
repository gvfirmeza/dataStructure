package main

import "fmt"

type fracao struct {
	num int
	den int
}

func adicionar(f1, f2 fracao) fracao {
	return fracao{num: f1.num*f2.den + f2.num*f1.den, den: f1.den * f2.den}
}

func subtrair(f1, f2 fracao) fracao {
	return fracao{num: f1.num*f2.den - f2.num*f1.den, den: f1.den * f2.den}
}

func multiplicar(f1, f2 fracao) fracao {
	return fracao{num: f1.num * f2.num, den: f1.den * f2.den}
}

func dividir(f1, f2 fracao) fracao {
	return fracao{num: f1.num * f2.den, den: f2.num * f1.den}
}

func mdc(a, b int) int {
	if b == 0 {
		return a
	}
	return mdc(b, a%b)
}

func simplificar(f fracao) fracao {
	if f.den == 0 {
		fmt.Println("Error: Division by zero")
		return fracao{num: 0, den: 1}
	}

	mdc := mdc(f.num, f.den)
	if f.num < 0 {
		mdc = -mdc
	}
	return fracao{num: f.num / mdc, den: f.den / mdc}
}


func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var f1, f2 fracao
		var operator rune
		fmt.Scan(&f1.num, &f1.den, &operator, &f2.num, &f2.den)

		var result fracao
		switch operator {
		case '+':
			result = adicionar(f1, f2)
		case '-':
			result = subtrair(f1, f2)
		case '*':
			result = multiplicar(f1, f2)
		case '/':
			result = dividir(f1, f2)
		}

		fmt.Printf("%d/%d = ", result.num, result.den)
		result = simplificar(result)
		fmt.Printf("%d/%d\n", result.num, result.den)
	}
}
