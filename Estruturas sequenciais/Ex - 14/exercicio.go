/* 14 - Faça um algoritmo que receba um valor A e B, e troque o valor de A por B e o valor de B por A e imprima na tela os valores.  */

package main

import (
	"fmt"
)

func main() {

	var a int

	fmt.Print("Digite o valor A: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Digite o valor B: ")
	fmt.Scan(&b)

	var temp int

	temp = a
	a = b
	b = temp

	fmt.Println("Após a troca os numero recebido: ")
	fmt.Print("O valor de A agora é: ", a, "\n")
	fmt.Print("O valor de B agora é: ", b, "\n")
}
