/*   11 - Faça um algoritmo que leia quatro notas obtidas por um aluno, calcule a média das nota obtidas, imprima na tela o nome do aluno e
se o aluno foi aprovado ou reprovado. Para o aluno ser considerado aprovado sua média final deve ser maior ou igual a 7.*/

package main

import "fmt"

 func main() {

	var name string

	fmt.Print("Digite seu nome:")
	fmt.Scan(&name)

	var nota1 int

	fmt.Print("Digite sua primeira nota:")
	fmt.Scan(&nota1)
	
	var nota2 int

	fmt.Print("Digite sua segunda nota:")
	fmt.Scan(&nota2)

	var nota3 int

	fmt.Print("Digite sua terceira nota:")
	fmt.Scan(&nota3)

	var nota4 int

	fmt.Print("Digite sua quarta nota:")
	fmt.Scan(&nota4)

   var notas int = int( nota1 + nota2 + nota3)

   resultado := notas / 4

	if resultado >= 7 {
		fmt.Print(resultado, " Parabéns ",name," você foi aprovado")
	}else {
		fmt.Print(resultado," Sinto muito ", name, " Você foi reprovado")
	}

 }	