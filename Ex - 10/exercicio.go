/*  10 - Faça um algoritmo que leia três notas obtidas por um aluno, e imprima na tela a média das notas. */

package main

import "fmt"

 func main() {

	var nota1 int

	fmt.Print("Digite sua primeira nota:")
	fmt.Scan(&nota1)
	
	var nota2 int

	fmt.Print("Digite sua segunda nota:")
	fmt.Scan(&nota2)

	var nota3 int

	fmt.Print("Digite sua terceira nota:")
	fmt.Scan(&nota3)

   var notas int = int( nota1 + nota2 + nota3)

   resultado := notas / 3

	fmt.Print("Sua média é: ",resultado)
 }	