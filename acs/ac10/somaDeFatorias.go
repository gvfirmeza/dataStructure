package main

import "fmt"

func funcFatorial(n int) uint64 {
	var fatorial uint64 = uint64(n)
	i := n - 1
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	}

	for i > 0 {
		fatorial = fatorial * uint64(i)
		i--
	}

	return fatorial
}

func main() {
	var input1, input2 int
	for {
		_, err := fmt.Scanf("%d", &input1)
		if err != nil {
			return
		}
		_, err = fmt.Scanf("%d", &input2)
		if err != nil {
			return
		}
		fatorial1 := funcFatorial(input1)
		fatorial2 := funcFatorial(input2)
		resultado := fatorial1 + fatorial2
		fmt.Println(resultado)
	}
}
