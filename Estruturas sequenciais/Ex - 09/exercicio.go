/*9 - Faça um algoritmo que calcule o IMC (Índice de Massa Corporal) de uma pessoa, leia o seu peso e sua altura e imprima na tela sua condição

de acordo com a tabela abaixo:

Fórmula do IMC = peso / (altura) ²

Tabela Condições IMC



 Abaixo de 18,5   | Abaixo do peso

 Entre 18,6 e 24,9 | Peso ideal (parabéns)

 Entre 25,0 e 29,9 | Levemente acima do peso

 Entre 30,0 e 34,9 | Obesidade grau I

 Entre 35,0 e 39,9 | Obesidade grau II (severa)

 Maior ou igual a 40 | Obesidade grau III (mórbida) */

package main

import "fmt"

 func main() {

	var peso float64
	var altura float64

	fmt.Print("Digite seu peso:")
	fmt.Scan(&peso)
	
	fmt.Print("Digite sua altura:")
	fmt.Scan(&altura)


	var resultado float64 =float64(peso / (altura * altura))

	if resultado < 18.6 {
		fmt.Print("IMC: ", resultado, " Abaixo do peso")
	}else if  resultado >= 18.6 && resultado < 29.9 {
		fmt.Print("IMC: ", resultado, " Peso ideal (parabéns)")
	}else if resultado >= 25.0 && resultado < 29.9 {
		fmt.Print("IMC: ", resultado, " Levemente acima do peso")
	}else if resultado >= 30.0 && resultado < 34.9 {
		fmt.Print("IMC: ", resultado, " Obesidade grau I ")
	}else if resultado >= 35.0 && resultado < 40  {
		fmt.Print("IMC: ", resultado, " Obesidade grau II (severa) ")
	}else {
		fmt.Print("IMC: ", resultado," Obesidade grau III (mórbida)")
	}



 }	