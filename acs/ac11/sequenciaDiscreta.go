package main

import "fmt"

func main() {
    var N int
    fmt.Scanf("%d", &N)

    sequencia := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d", &sequencia[i])
    }

    var consecutivos int
    var atual int
    for i := 0; i < N; i++ {
        if sequencia[i] != atual {
            consecutivos++
        }
        atual = sequencia[i]
    }

    fmt.Println(consecutivos)
}

