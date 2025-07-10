/*
Faça um programa para uma loja de tintas. O programa deverá pedir o tamanho em metros quadrados da área a ser pintada.

Considere que a cobertura de tinta é de 1 litro para cada 3 metros quadrados e que a tinta é vendida em latas de 18 litros, que custa R$ 80,00.

Informe ao usuário as quantidades de latas de tinta que serão compradas e o preço total.
*/

package main

import (
	"fmt"
)

func main() {

	var area float64
	fmt.Print("Digite o tamanho da área a ser pintada em metros quadrados: ")
	fmt.Scanln(&area)

	const coberturaPorLitro = 3.0
	const litrosPorLata = 18.0
	const precoPorLata = 80.0

	litrosNecessarios := area / coberturaPorLitro // quantidade de litros necessários

	// quantidade de latas necessárias (arredondando para cima)
	latasNecessarias := int(litrosNecessarios/litrosPorLata + 0.9999) // Arredondamento para cima

	// Calculo do preço total
	precoTotal := float64(latasNecessarias) * precoPorLata

	fmt.Printf("Quantidade de latas de tinta necessárias: %d\n", latasNecessarias)
	fmt.Printf("Preço total: R$ %.2f\n", precoTotal)

}
