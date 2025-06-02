//4 - Faça um algoritmo que receba um número inteiro e imprima na tela o seu antecessor e o seu sucessor.

package main

import (
	"fmt"
)

func main() {
var a int 
fmt.Print("Digite um número: ")
fmt.Scan(&a)
 
b := a - 1
c := a + 1

fmt.Print(" o antecessor de: ", a, " é ", b , " e o sucessor dele é: ", c)
 
}