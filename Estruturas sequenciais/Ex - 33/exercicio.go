/*
Faça um programa que peça o tamanho de um arquivo para download
(em MB) e a velocidade de um link de Internet (em Mbps),
 calcule e informe o tempo aproximado de download do arquivo usando este link (em minutos).
*/

package main

import (
	"fmt"
)

func main() {
	var tamanhoArquivoMb float64
	var velocidadeLinkMbps float64

	fmt.Print("Digite o tamanho do arquivo para download (em MB): ")
	fmt.Scanln(&tamanhoArquivoMb)

	fmt.Print("Digite a velocidade do link de Internet (em Mbps): ")
	fmt.Scanln(&velocidadeLinkMbps)

	// Convertendo Mbps para MBps (Megabytes por segundo)
	velocidadeLinkMBps := velocidadeLinkMbps / 8.0

	// Calculando o tempo de download em segundos
	tempoDownloadSegundos := tamanhoArquivoMb / velocidadeLinkMBps

	// Convertendo o tempo de download para minutos
	tempoDownloadMinutos := tempoDownloadSegundos / 60.0
	fmt.Printf("Tempo aproximado de download: %.2f minutos\n", tempoDownloadMinutos)

}
