/* 18 - Francisco tem 1,50m e cresce 2 centímetros por ano, enquanto Sara tem 1,10m e cresce 3 centímetros por ano.
Faça um algoritmo que calcule e imprima na tela em quantos anos serão necessários para que Francisco seja maior que Sara. */

package main

import (
	"fmt"
)

func main() {
	alturaFrancisco := 1.50
	alturaSara := 1.10
	crescimentoFrancisco := 0.02
	crescimentoSara := 0.03
	anos := 0

	if alturaFrancisco > alturaSara {
		fmt.Println("Francisco já é maior que Sara.")
	} else {
		for alturaFrancisco <= alturaSara {
			alturaFrancisco += crescimentoFrancisco
			alturaSara += crescimentoSara
			anos++
		}
		fmt.Println("Francisco será maior que Sara em", anos, "anos.")
	}

	fmt.Print("francisco foi maior em ", anos, " anos")
}
