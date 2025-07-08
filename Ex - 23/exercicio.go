/*  Faça um algoritmo que efetue o cálculo do salário líquido de um professor.
As informações fornecidas serão: valor da hora aula,
número de aulas lecionadas no mês e percentual de desconto do INSS.
Imprima na tela o salário líquido final.*/

package main

import (
	"fmt"
)

func main() {
	var valorHora, numeroAulas, percentualDesconto float64
	fmt.Print("Digite o valor da hora aula:")
	fmt.Scanln(&valorHora)

	fmt.Print("Digite o número de aulas lecioonadas no mês: ")
	fmt.Scanln(&numeroAulas)

	fmt.Print("Digite o percentual de desconto do INSS:")
	fmt.Scanln(&percentualDesconto)

	salarioBruto := valorHora * numeroAulas
	desconto := (percentualDesconto / 100) * salarioBruto
	salarioBruto -= desconto
	fmt.Printf("O salário líquido do professor é: R$ %.2f\n", salarioBruto)

}
