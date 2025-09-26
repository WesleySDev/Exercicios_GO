/*
Faça um programa que pergunte o preço de três produtos e informe qual produto você deve comprar, sabendo que a decisão é sempre pelo mais barata.
*/
package main

import (
	"fmt"
)

func main() {
	var preco1, preco2, preco3 float64

	fmt.Print("Digite o preço do primeiro produto: ")
	fmt.Scan(&preco1)

	fmt.Print("Digite o preço do segundo produto: ")
	fmt.Scan(&preco2)

	fmt.Print("Digite o preço do terceiro produto: ")
	fmt.Scan(&preco3)

	if preco1 < preco2 && preco1 < preco3 {
		fmt.Println("Você deve comprar o primeiro produto.")
	} else if preco2 < preco1 && preco2 < preco3 {
		fmt.Println("Você deve comprar o segundo produto.")
	} else {
		fmt.Println("Você deve comprar o terceiro produto.")
	}
}
