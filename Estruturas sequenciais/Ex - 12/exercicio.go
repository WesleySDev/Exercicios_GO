/*  12 - Faça um algoritmo que leia o valor de um produto e determine o valor que deve ser pago, conforme a escolha da forma de pagamento

pelo comprador e imprima na tela o valor final do produto a ser pago. Utilize os códigos da tabela de condições de pagamento para efetuar o cálculo adequado.



Tabela de Código de Condições de Pagamento



1 - À Vista em Dinheiro ou Pix, recebe 15% de desconto

2 - À Vista no cartão de crédito, recebe 10% de desconto

3 - Parcelado no cartão em duas vezes, preço normal do produto sem juros

4 - Parcelado no cartão em três vezes ou mais, preço normal do produto mais juros de 10% */

package main

import "fmt"

 func main() {

	var preco_produto float64
	Dinheiro_pix := 1
	cartao_vista := 2
	Parcelado_em_dois := 3
	Parcelado_em_varias := 4

	fmt.Print("preço do produto? ")
	fmt.Scan(&preco_produto)


	var forma_pagamento int
	fmt.Print("Forma de pagamento?\n1 - Dinheiro ou Pix\n2 - Cartão à vista\n3 - Parcelado em 2x\n4 - Parcelado em várias vezes\n")
	fmt.Scan(&forma_pagamento)




  if forma_pagamento == Dinheiro_pix  {
	fmt.Print(preco_produto * (1 - 0.15) )
  }else if forma_pagamento ==  cartao_vista {
	fmt.Print(preco_produto * (1 - 0.1) )
  } else if forma_pagamento == Parcelado_em_dois {
		fmt.Println("Você escolheu parcelar em 2 vezes.")
		fmt.Println("Valor de cada parcela: ", preco_produto/2)
 }else if forma_pagamento == Parcelado_em_varias  {
	var num_parcelas int
	fmt.Print("Em quantas vezes deseja parcelar? ")
	fmt.Scan(&num_parcelas)

	if num_parcelas < 3 {
		fmt.Print("Número de parcelas inválido. Para esta opção, deve ser 3 ou mais.")
	}else {
		preco_com_juros := preco_produto * 1.10
		valor_parcela := preco_com_juros / float64(num_parcelas)
		
		fmt.Println("Você escolheu parcelar em", num_parcelas, "vezes.")
		fmt.Println("Preço total com juros: ", preco_com_juros)
		fmt.Println("Valor de cada parcela: ", valor_parcela)
	}
 }else {
	fmt.Println("Forma de pagamento inválida.")
 }
}
 