//3 - Faça um algoritmo que leia dois valores inteiros A e B, se os valores de A e B forem iguais, deverá somar os dois valores,

//caso contrário devera multiplicar A por B. Ao final de qualquer um dos cálculos deve-se atribuir o resultado a uma variável C e

//imprimir seu valor na tela.

package main

import (
	"fmt"
)

func main() {
	var a, b int
	var c int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)

	if a == b {
		c = a + b
		fmt.Print(c)
	} else {
		c = a * b
		fmt.Print(c)
	}

}
