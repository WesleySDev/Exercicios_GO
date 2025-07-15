//5 - Faça um algoritmo que leia o valor do salário mínimo e o valor do salário de um usuário, calcule quantos salários mínimos esse

//usuário ganha e imprima na tela o resultado. (Base para o Salário mínimo R$ 1.293,20).

package main

import (
	"fmt"
)

func main() {
a := 1293.20

var salarioUsuario float64
fmt.Print("Digite seu sálario: ")
fmt.Scan(&salarioUsuario)
var resultado int = int(salarioUsuario / a)

fmt.Print("Você recebe: ", resultado, " salários mínimo")

}