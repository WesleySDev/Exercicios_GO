/* 16 - Faça um algoritmo que leia três valores que representam os três lados de um triângulo e verifique se são válidos, determine se o triângulo é

equilátero, isósceles ou escaleno. */

package main

import (
	"fmt"
)

func main() {
	var lado1, lado2, lado3 int

	fmt.Print("Digite o primero lado do seu triangulo: ")
	fmt.Scan(&lado1)
	fmt.Print("Digite o segundo lado do seu triangulo: ")
	fmt.Scan(&lado2)
	fmt.Print("Digite o terceiro lado do seu triangulo: ")
	fmt.Scan(&lado3)

	if (lado1+lado2 > lado3) && (lado1+lado3 > lado2) && (lado2+lado3 > lado1) {
		if lado1 == lado2 && lado2 == lado3 {
			fmt.Print("Seu tringule é um Equilátero")
		} else if lado1 == lado2 || lado1 == lado3 || lado2 == lado3 {
			fmt.Print("Seu tringule é um Isósceles")
		} else if lado1 != lado2 && lado1 != lado3 && lado2 != lado3 {
			fmt.Print("Seu tringule é um Escaleno")
		}
	} else {
		fmt.Print("Nao é um triângulo válido")
	}

}
