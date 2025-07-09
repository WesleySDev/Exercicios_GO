/*
Faça um Programa que parte o raio de um círculo, calcule e mostre sua área.
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var raio float64
	fmt.Print("Digite o raio do circulo: ")
	fmt.Scanln(&raio)

	area := math.Pi * math.Pow(raio, 2)
	fmt.Printf("A área do círculo com raio %.2f é: %.2f\n", raio, area)
}
