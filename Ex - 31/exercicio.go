/*
João Papo-de-Pescador, homem de bem, comprou um microcomputador para controlar o rendimento diário de seu trabalho.

Toda vez que ele trouxer um peso de peixes maior que o previsto pelo regulamento de pesca do estado de São Paulo (50 quilos) deverá pagar uma multa de R$ 4,00 por quilo excedente.

João precisa que você faça um programa que leia a variável peso (peso de peixes) e calcule o excesso.

Gravar na variável excesso a quantidade de quilos além do limite e na multa variável o valor da multa que João deverá pagar. Imprima os dados do programa com as mensagens apropriadas.
*/

package main

import (
	"fmt"
)

func main() {

	var peso float64
	fmt.Print("Digite o peso dos peixes pescados (em quilos): ")
	fmt.Scanln(&peso)

	const limite float64 = 50.0
	const multaPorQuilo float64 = 4.0

	if peso > limite {
		excesso := peso - limite
		multa := excesso * multaPorQuilo
		fmt.Printf("Excesso de peso: %.2f kg\n", excesso)
		fmt.Printf("Valor da multa a ser paga: R$ %.2f\n", multa)
	} else {
		fmt.Println("Não há excesso de peso. Nenhuma multa a ser paga.")
	}

}
