/*
Faça um programa para a leitura de duas notas parciais de um aluno.
O programa deve calcular a média alcançada por aluno e apresentar:

	A mensagem "Aprovado", se a média alcançada for maior ou igual a sete; A mensagem "Reprovado",
	se a mídia for menor do que sete; A mensagem "Aprovado com Distinção", se a média for igual a dez.
*/
package main

import (
	"fmt"
)

func main() {
	var nota1 float64
	var nota2 float64

	soma := nota1 + nota2

	fmt.Print("Digite a primera nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	media := soma / 2

	if media >= 7 {
		fmt.Println(media, " Aprovado")
	} else if media < 7 {
		fmt.Println(media, " Reprovado")
	} else if media == 10 {
		fmt.Println(media, " Aprovado com Distinção")
	}

}
