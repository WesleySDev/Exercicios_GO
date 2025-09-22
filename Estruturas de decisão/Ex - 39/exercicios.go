/*
Faça um Programa que leia três números e mostre o maior deles
*/
package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var num3 float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Println("O maior número é: ", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("O maior número é: ", num2)
	} else {
		fmt.Println("O maior número é: ", num3)
	}
}
