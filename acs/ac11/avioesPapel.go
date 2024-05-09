package main

import "fmt"

func main() {
    var c, p, f int
    fmt.Scanf("%d %d %d", &c, &p, &f)

    if p >= c*f {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}