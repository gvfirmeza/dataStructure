package main

import "fmt"

func mdc(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    var n int
    fmt.Scanf("%d", &n)
    for i := 0; i < n; i++ {
        var x, y int
        var op string
        fmt.Scanf("%d %s %d", &x, &op, &y)
        var res int
        switch op {
        case "+":
            res = x*y + y*x
        case "-":
            res = x*y - y*x
        case "*":
            res = x * y
        case "/":
            res = x * y
        }
        num, den := res, x*y
        gcd := mdc(num, den)
        fmt.Printf("%d/%d\n", num/gcd, den/gcd)
    }
}
