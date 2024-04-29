package main

import (
	"fmt"
	"math"
)

func primo(num int64) bool {
	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	maxDiv := int64(math.Sqrt(float64(num))) + 1
	for i := int64(2); i <= maxDiv; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var c, numc int
	fmt.Scan(&c)

	for numc < c {
		var num int64
		fmt.Scan(&num)

		if primo(num) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}

		numc++
	}
}
