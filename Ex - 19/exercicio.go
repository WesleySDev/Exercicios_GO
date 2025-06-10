/* 19 - Faça um algoritmo que imprima na tela a tabuada de 1 até 10. */

package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Tabuada do", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}

	}
}
