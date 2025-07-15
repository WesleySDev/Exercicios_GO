/* 21 - Faça um algoritmo que mostre um valor aleatório entre 0 e 100. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())         // Inicializa a "semente" com o tempo atual
	numero := rand.Intn(101)                 // Gera número aleatório de 0 até 100 (101 não incluso)
	fmt.Println("Número aleatório:", numero) // Mostra o número
}
