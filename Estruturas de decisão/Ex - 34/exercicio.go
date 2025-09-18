/*
Faça um Programa que peça dois números e imprima o maior deles.
*/

package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Printf("O maior número é: %d\n", num1)
	} else if num2 > num1 {
		fmt.Printf("O maior número é: %d\n", num2)
	} else {
		fmt.Println("Os números são iguais.")
	}
}
