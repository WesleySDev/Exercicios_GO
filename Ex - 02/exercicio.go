//2 - Faça um algoritmo para receber um número qualquer e imprimir na tela se o número é par ou ímpar, positivo ou negativo.

package main

import (
	"fmt"
)

func main() {

	var NumeroColetado int 
   var par_impar string
   var positivo_neagtivo string

   
	fmt.Print("digite um numero: ")
	fmt.Scan(&NumeroColetado)

   if NumeroColetado % 2 == 0 {
    par_impar = "par"
   }else  {
     par_impar = "impar"
   }
   
   if NumeroColetado > 0  {
      positivo_neagtivo = "positivo"
   }else if NumeroColetado < 0 {
      positivo_neagtivo = "nagativo"
   }else {
      positivo_neagtivo = "zero"
   }
   fmt.Print("O numero: ", NumeroColetado, " é ", par_impar , " e ", positivo_neagtivo )

}