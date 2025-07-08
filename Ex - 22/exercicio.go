/* 22 - Faça um algoritmo que leia dois valores inteiros A e B,
imprima na tela o quociente e o resto da divisão inteira entre eles. */

package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)

	divisao := a / b
	resto := a % b

	fmt.Print("Resto: ", resto, "\n")
	fmt.Print("Quociente: ", divisao, "\n")
}
