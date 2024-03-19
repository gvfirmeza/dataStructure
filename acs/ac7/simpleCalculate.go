package main
 
import (
    "fmt"
)
 
func main() {
    var a, b int 
    var c float64

	fmt.Scanln(&a, &b, &c)
	
	var d, e int 
    var f float64

	fmt.Scanln(&d, &e, &f)
	
    resultado := float64(b)*c + float64(e)*f
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", resultado)
}