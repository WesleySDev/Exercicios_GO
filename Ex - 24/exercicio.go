/*
Faça um algoritmo que calcule a quantidade de litros de combustível gastos em uma viagem, sabendo que o carro faz 12km com um litro.
 Deve-se fornecer ao usuário o tempo que será gasto na viagem a sua
 velocidade média, distância percorrida e a
 quantidade de litros utilizados para fazer a viagem*/

package main

import (
	"fmt"
)

func main() {

	var tempoViagem, velocidadeMedia, distanciaPercorrida, litrosGastos float64
	fmt.Print("Digite o tempo da viagem em horas: ")
	fmt.Scanln(&tempoViagem)

	fmt.Print("Digite a velocidade média em km/h: ")
	fmt.Scanln(&velocidadeMedia)

	fmt.Print("Digite a distância percorrida em km: ")
	fmt.Scanln(&distanciaPercorrida)

	distanciaPercorrida = velocidadeMedia * tempoViagem

	litrosGastos = distanciaPercorrida / 12.0

	fmt.Printf("O tempo gasto na viagem é: %.2f horas\n", tempoViagem)
	fmt.Printf("A velocidade média da viagem é: %.2f km/h\n", velocidadeMedia)
	fmt.Printf("A distância percorrida na viagem é: %.2f km\n", distanciaPercorrida)
	fmt.Printf("A quantidade de litros utilizados para fazer a viagem é: %.2f litros\n", litrosGastos)
}
