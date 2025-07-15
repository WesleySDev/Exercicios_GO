//8 - Faça um algoritmo que leia três valores inteiros diferentes e imprima na tela os valores em ordem decrescente.

package main

import (
	"fmt"
	"sort"
)

func main() {

 var valores [3]int

	fmt.Print("Digite o primeiro valor:")
	fmt.Scan(&valores[0])

	fmt.Print("Digite o primeiro valor:")
	fmt.Scan(&valores[1])

	fmt.Print("Digite o primeiro valor:")
	fmt.Scan(&valores[2])

	sort.Ints(valores[:])
  	
	fmt.Print(valores[2], valores[1], valores[0])
	} 	