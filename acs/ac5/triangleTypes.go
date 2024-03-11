package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	fmt.Scanln(&a, &b, &c)

	if a < b {
		a, b = b, a
	}
	if b < c {
		b, c = c, b
	}
	if a < b {
		a, b = b, a
	}

	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if a*a == b*b+c*c {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if a*a > b*b+c*c {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if a*a < b*b+c*c {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if a == b && b == c {
			fmt.Println("TRIANGULO EQUILATERO")
		}
		if (a == b && a != c) || (a == c && a != b) || (b == c && b != a) {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}
}
