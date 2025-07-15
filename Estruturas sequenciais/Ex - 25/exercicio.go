/*
Faça um Programa que converte metros para centímetros.
*/

package main

import (
	"fmt"
)

func main() {
	var valorMetros int

	fmt.Print("Digite o valor em metros: ")
	fmt.Scanln(&valorMetros)

	valorTotalCentimetros := (valorMetros * 100)

	fmt.Printf("O valor total em metros é: %d m\n", valorMetros)
	fmt.Printf("O valor total em centímetros é: %d cm\n", valorTotalCentimetros)

}
