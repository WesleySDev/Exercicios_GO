/*20 - Faça um algoritmo que receba um valor inteiro e imprima na tela a sua tabuada. */

package main

import (
	"fmt"
)

func main() {

	var i int
	fmt.Print("Digite um valor: ")
	fmt.Scan(&i)
	fmt.Println("Tabuada do", i)
	for j := 1; j <= 10; j++ {
		fmt.Printf("%d x %d = %d\n", i, j, i*j)
	}

}
