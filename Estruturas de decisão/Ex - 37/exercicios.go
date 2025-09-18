/*
Faça um Programa que verifique se uma letra digitada é vogal ou consoante.
*/
package main

import (
	"fmt"
)

func main() {
	var letra string

	fmt.Print("Digite uma letra: ")
	fmt.Scan(&letra)

	if letra == "A" || letra == "E" || letra == "I" || letra == "O" || letra == "U" ||
		letra == "a" || letra == "e" || letra == "i" || letra == "o" || letra == "u" {
		fmt.Print(letra, " Essa letra é vogal")
	} else {
		fmt.Print(letra, " Essa letra é consoante")
	}
}
