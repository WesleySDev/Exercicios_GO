/* 17 - Faça um algoritmo que leia uma temperatura em Fahrenheit e calcule a temperatura correspondente em grau Celsius. Imprima na tela as duas temperaturas.

Fórmula: C = (5 * ( F-32) / 9) */

package main

import (
	"fmt"
)

func main() {
	var fahrenheit int

	fmt.Print("Digite A temperatura em Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (5 * (fahrenheit - 32) / 9)

	fmt.Print("A temperatura em Fahrenheit é: ", fahrenheit, "\nA temperatura em grau Celsius é: ", celsius)

}
