/*
Faça um Programa que leia três números e mostre-os em ordem decrescente.
*/
package main

import (
	"fmt"
)

func main() {

	var num1, num2, num3 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)
	if num1 >= num2 && num1 >= num3 {
		if num2 >= num3 {
			fmt.Printf("Ordem decrescente: %.2f, %.2f, %.2f\n", num1, num2, num3)
		} else {
			fmt.Printf("Ordem decrescente: %.2f, %.2f, %.2f\n", num1, num3, num2)
		}
	} else if num2 >= num1 && num2 >= num3 {
		if num1 >= num3 {
			fmt.Printf("Ordem decrescente: %.2f, %.2f, %.2f\n", num2, num1, num3)
		} else {
			fmt.Printf("Ordem decrescente: %.2f, %.2f, %.2f\n", num2, num3, num1)
		}
	} else {
		if num1 >= num2 {
			fmt.Printf("Ordem decrescente: %.2f, %.2f, %.2f\n", num3, num1, num2)
		} else {
			fmt.Printf("Ordem decrescente: %.2f, %.2f, %.2f\n", num3, num2, num1)
		}
	}

}
