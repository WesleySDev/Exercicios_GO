/* 13 - Faça algoritmo que leia o nome e a idade de uma peso e imprima na tela o nome da pessoa e se ela é maior ou menor de idade. */

package main

import "fmt"

 func main() {
    var name string

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&name)

	var idade int

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	if idade >= 18 {
	fmt.Print("Parabéns ", name, " você é maior de idade")
	}else {
		fmt.Print("Infelizmente ", name, " você é menor de idade")
	}
}
