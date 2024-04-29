package main

import (
    "fmt"
    "math"
)

func main() {
    var X1, Y1, X2, Y2 int
    saida := 2
    var x, y int

    for {
        n, err := fmt.Scanf("%d %d %d %d", &X1, &Y1, &X2, &Y2)
        if n != 4 || err != nil || (X1 == 0 && Y1 == 0 && X2 == 0 && Y2 == 0) {
            break
        }

        x = int(math.Abs(float64(X1 - X2)))
        y = int(math.Abs(float64(Y1 - Y2)))

        if X1 == X2 && Y1 == Y2 {
            saida = 0
            fmt.Println(saida)
        } else if X1 == X2 && Y1 != Y2 {
            saida = 1
            fmt.Println(saida)
        } else if X1 != X2 && Y1 == Y2 {
            saida = 1
            fmt.Println(saida)
        } else if x == y {
            saida = 1
            fmt.Println(saida)
        } else {
            saida = 2
            fmt.Println(saida)
        }
    }
}

