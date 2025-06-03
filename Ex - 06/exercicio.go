//6 - Faça um algoritmo que leia um valor qualquer e imprima na tela com um reajuste de 5%.

package main

import (
	"fmt"
)

func main() {

var valor float64

fmt.Print("Digite um valor: ")
fmt.Scan(&valor)

r := 1.05
var resultado float64 = float64( valor * float64(r)) 

fmt.Print(resultado)
} 	