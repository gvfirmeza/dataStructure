package main

import "fmt"

func main() {
    var n, dias int
    var c float64

    fmt.Scanf("%d", &n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%f", &c)
        dias = 0
        for c > 1 {
            c /= 2
            dias++
        }
        fmt.Printf("%d dias\n", dias)
    }
}