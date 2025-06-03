//7 - Faça um algoritmo que leia dois valores booleanos (lógicos) e determine se ambos são VERDADEIRO ou FALSO.

package main

import (
	"fmt"
)

func main() {

var primeiro_valor bool
var segundo_valor bool

fmt.Print("Digite o primeiro valor ( true ou falso ): \n")
fmt.Scan(&primeiro_valor)

fmt.Print("Digite o Segundo valor ( true ou falso ): \n")
fmt.Scan(&segundo_valor)

if  primeiro_valor == true && segundo_valor == true {
  fmt.Print("ambos verdadeiro\n")

}else if primeiro_valor == false && segundo_valor == false {
	fmt.Print("ambos falso\n")
}else  {
	fmt.Print("um verdadeiro e outro falso")
}
	} 	