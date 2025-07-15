/* 15 - Faça um algoritmo que leia o ano em que uma pessoa nasceu, imprima na tela quantos anos, meses e dias essa pessoa ja viveu. Leve em

consideração o ano com 365 dias e o mês com 30 dias.

(Ex: 5 anos, 2 meses e 15 dias de vida) */

package main

import (
	"fmt"
)

func main() {
	var anoNasc, mesNasc, diaNasc int
	var anoAtual, mesAtual, diaAtual int

	fmt.Print("Digite o ano de seu nacimento: ")
	fmt.Scan(&anoNasc)

	fmt.Print("Digite o mês de seu nacimento: ")
	fmt.Scan(&mesNasc)

	fmt.Print("Digite o dia de seu nacimento: ")
	fmt.Scan(&diaNasc)

	fmt.Print("Digite o ano atual: ")
	fmt.Scan(&anoAtual)

	fmt.Print("Digite o mês atual: ")
	fmt.Scan(&mesAtual)

	fmt.Print("Digite o dia atual: ")
	fmt.Scan(&diaAtual)

	diasDeNacimento := (anoNasc * 365) + (mesNasc * 30) + diaNasc
	diasAtuais := (anoAtual * 365) + (mesAtual + 30) + diaAtual

	totalDias := diasAtuais - diasDeNacimento

	anos := totalDias / 365
	meses := (totalDias % 365) / 30
	dias := (totalDias % 365) % 30

	fmt.Printf("Voce viveu aproximadamente: %d anos, %d meses e %d dias.\n", anos, meses, dias)
}
