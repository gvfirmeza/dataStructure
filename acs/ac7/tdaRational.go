package main

import (
  "fmt"
)

func mdc(a, b int) int {
  for b != 0 {
    a, b = b, a%b
  }
  return a
}

func simplificar(numerador, denominador int) (int, int) {
  divisor := mdc(numerador, denominador)
  return numerador / divisor, denominador / divisor
}

func calcular(numerador1, denominador1, numerador2, denominador2 int, operacao string) (int, int, error) {
  switch operacao {
  case "+":
    return numerador1*denominador2 + numerador2*denominador1, denominador1 * denominador2, nil
  case "-":
    return numerador1*denominador2 - numerador2*denominador1, denominador1 * denominador2, nil
  case "*":
    return numerador1 * numerador2, denominador1 * denominador2, nil
  case "/":
    if denominador2 == 0 {
      return 0, 0, fmt.Errorf("divisão por zero")
    }
    return numerador1 * denominador2, numerador2 * denominador1, nil
  default:
    return 0, 0, fmt.Errorf("operação inválida: %s", operacao)
  }
}

func main() {
  var numeroCasos int
  fmt.Scanf("%d", &numeroCasos)

  for i := 0; i < numeroCasos; i++ {
    var numerador1, denominador1, numerador2, denominador2 int
    var operacao string
    fmt.Scanf("%d/%d %s %d/%d", &numerador1, &denominador1, &operacao, &numerador2, &denominador2)

    resultadoNumerador, resultadoDenominador, err := calcular(numerador1, denominador1, numerador2, denominador2, operacao)
    if err != nil {
      fmt.Println("Erro:", err.Error())
      continue
    }

    numeradorSimplificado, denominadorSimplificado := simplificar(resultadoNumerador, resultadoDenominador)
    fmt.Printf("%d/%d %s %d/%d = %d/%d\n", numerador1, denominador1, operacao, numerador2, denominador2, numeradorSimplificado, denominadorSimplificado)
  }
}