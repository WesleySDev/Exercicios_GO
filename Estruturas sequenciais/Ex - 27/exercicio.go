/*
Faça um programa que calcula a área de um quadrado, em seguida mostre o dobro desta área para o usuário.
*/

package main

import (
	"fmt"
)

func main() {

	var ladoQuadrado float64
	fmt.Print("Digite o lado do quadrado: ")
	fmt.Scanln(&ladoQuadrado)

	areaQuadrado := ladoQuadrado * ladoQuadrado
	dobroArea := areaQuadrado * 2

	fmt.Printf("A área do quadrado é: %.2f\n", areaQuadrado)
	fmt.Printf("O dobro da área do quadrado é: %.2f\n", dobroArea)

}
