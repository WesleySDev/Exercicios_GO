/*
Faça um Programa que peça 2 números inteiros e um número real.
 Calcule e mostre: o produto do dobro do primeiro com metade do segundo.
  a soma do triplo do primeiro com o terceiro. o terceiro elevado ao cubo.
*/

package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	var num3 float64

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scanln(&num2)
	fmt.Print("Digite um número real: ")
	fmt.Scanln(&num3)

	produto := 2 * num1 * (num2 / 2)
	soma := 3*float64(num1) + num3
	terceiroElevado := num3 * num3 * num3

	fmt.Printf("O produto do dobro do primeiro com metade do segundo é: %d\n", produto)
	fmt.Printf("A soma do triplo do primeiro com o terceiro é: %.2f\n", soma)
	fmt.Printf("O terceiro elevado ao cubo é: %.2f\n", terceiroElevado)

}
