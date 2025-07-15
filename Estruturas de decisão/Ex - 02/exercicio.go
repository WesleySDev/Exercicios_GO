/*
Faça um Programa que peça um valor e mostre na tela se o valor é positivo ou negativo.
*/

package main

import "fmt"

func main() {

	var valor float64

	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)

	if valor > 0 {
		fmt.Println("O valor é positivo.")
	} else if valor < 0 {
		fmt.Println("O valor é negativo.")
	} else {
		fmt.Println("O valor é zero.")
	}
}
