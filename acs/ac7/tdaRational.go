package main

import "fmt"

type Fracao struct {
    num int
    den int
}

func soma(f1 Fracao, f2 Fracao) Fracao {
    retorno := Fracao{}
    retorno.num = f1.num*f2.den + f2.num*f1.den
    retorno.den = f1.den * f2.den
    return retorno
}

func subtracao(f1 Fracao, f2 Fracao) Fracao {
    retorno := Fracao{}
    retorno.num = f1.num*f2.den - f2.num*f1.den
    retorno.den = f1.den * f2.den
    return retorno
}

func multiplicacao(f1 Fracao, f2 Fracao) Fracao {
    retorno := Fracao{}
    retorno.num = f1.num * f2.num
    retorno.den = f1.den * f2.den
    return retorno
}

func divisao(f1 Fracao, f2 Fracao) Fracao {
    retorno := Fracao{}
    retorno.num = f1.num * f2.den
    retorno.den = f2.num * f1.den
    return retorno
}

func MDC(a int, b int) int {
    if b == 0 {
        return a
    }
    return MDC(b, a%b)
}

func irredutivel(f Fracao) Fracao {
    mdc := 0
    if f.num < 0 {
        mdc = MDC(-f.num, f.den)
    } else {
        mdc = MDC(f.num, f.den)
    }
    f.num /= mdc
    f.den /= mdc
    return f
}

func main() {
    var N int
    var operador byte
    var f1 Fracao
    var f2 Fracao
    var resultado Fracao
    fmt.Scanf("%d\n", &N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d / %d %c %d / %d\n", &f1.num, &f1.den, &operador, &f2.num, &f2.den)
        switch operador {
        case '+':
            resultado = soma(f1, f2)
            break
        case '-':
            resultado = subtracao(f1, f2)
            break
        case '*':
            resultado = multiplicacao(f1, f2)
            break
        case '/':
            resultado = divisao(f1, f2)
            break
        }
        fmt.Printf("%d/%d = ", resultado.num, resultado.den)
        resultado = irredutivel(resultado)
        fmt.Printf("%d/%d\n", resultado.num, resultado.den)
    }
}

