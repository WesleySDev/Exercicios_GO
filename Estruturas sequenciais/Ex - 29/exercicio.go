/*
Tendo como dados de entrada a altura de uma pessoa, construa um algoritmo que calcula seu peso ideal,
 usando a seguinte fórmula: (72.7*altura) - 58
*/

package main

import (
	"fmt"
)

func main() {

	var altura float64
	fmt.Print("Digite sua altura: ")
	fmt.Scanln(&altura)
	pesoIdeal := (72.7 * altura) - 58
	fmt.Printf("O peso ideal para uma altura de %.2f metros é: %.2f kg\n", altura, pesoIdeal)
}
