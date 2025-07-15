/*
Tendo como dado de entrada a altura (h) de uma pessoa,
 constrói um algoritmo que calcula seu peso ideal,
  utilizando as seguintes fórmulas: Para homens: (72,7 h) - 58 Para mulheres: (62,1 h) - 44,7
*/

package main

import (
	"fmt"
)

func main() {

	var altura float64
	var sexo string
	fmt.Print("Digite sua altura: ")
	fmt.Scanln(&altura)
	fmt.Print("Digite seu sexo (M para masculino, F para feminino): ")
	fmt.Scanln(&sexo)
	var pesoIdeal float64
	if sexo == "M" || sexo == "m" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" || sexo == "f" {
		pesoIdeal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Sexo inválido. Por favor, digite 'M' para masculino ou 'F' para feminino.")
	}
	fmt.Printf("O peso ideal para uma altura de %.2f metros é: %.2f kg\n", altura, pesoIdeal)
}
