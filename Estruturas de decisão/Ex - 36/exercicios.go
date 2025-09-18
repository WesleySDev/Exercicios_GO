/*
	Faça um Programa que verifique se uma letra digitada é "F" ou "M".

Conforme a letra escrever: F - Feminino, M - Masculino, Sexo Inválido
*/
package main

import (
	"fmt"
)

func main() {

	var valor string

	fmt.Print("Digite uma letra: ")
	fmt.Scan(&valor)

	if valor == "M" || valor == "m" {
		fmt.Println(valor, " Sexo Mascculino ")
	} else if valor == "F" || valor == "f" {
		fmt.Println(valor, " Sexo Feminino")
	} else {
		fmt.Println("Sexo invalido tente: F - Feminino, M - Masculino ")
	}

}
