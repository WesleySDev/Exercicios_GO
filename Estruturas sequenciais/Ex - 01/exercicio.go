//1 - Faça um algoritmo que leia os valores de A, B, C e em seguida imprima na tela a soma entre A e B é mostre se a soma é menor que C.

package main

import "fmt"

func main() {
  a := 10
  fmt.Print(" 1 variavel:", a,"\n")
  b := 20
	fmt.Print(" 2 variavel - ", b,"\n")
  c := 30
	fmt.Print(" 3 variavel - ", c,"\n")

    d := a + b

   if d < c {
	fmt.Print(" a soma é: ", d, " por isso é menor que: ", c,"\n")

   }else if d > c{
	fmt.Print(" a soma é: ", d, " por isso é maior que: ", c,"\n")
   }else {
	fmt.Print(" os numeros ", d, " e ", c, " são indenticos.\n")
   } 


}