package main

import (
	"fmt"
)

func main() {
	// DECLARAÇÃO DE VARIAVEL COM ARGUMENTOS:
	var nome string = "Willian Alves Fonseca"
	fmt.Println(nome)
	var idade int = 28
	fmt.Println("A minha idade é: ", idade)
	var altura float32 = 1.73
	fmt.Println("A minha altura é: ", altura)

	fmt.Println("======================")
	/*DECLARANDO VARIAVEIS DE FORMA RESUMIDA:
	Neste caso a variavel coloca implicitamente
	o tipo de dado de acordo com o valor atribuido.
	*/
	var nome2 = "Jacinto a Misericordia de Deus"
	fmt.Println(nome2)
	var idade2 = 33
	fmt.Println("A minha idade é: ", idade2)
	var altura2 = 1.80
	fmt.Println("A minha altura é: ", altura2)

	fmt.Println("======================")
	/*DECLARANDO VARIAVEIS DE FORMA SUPER RESUMIDA:
	Neste caso a variavel coloca implicitamente
	o tipo de dado de acordo com o valor atribuido.
	*/
	nome3 := "Willian Gostosão da Bala Chita"
	fmt.Println(nome3)
	idade3 := 25
	fmt.Println("A minha idade é: ", idade3)
	altura3 := 1.90
	fmt.Println("A minha altura é: ", altura3)
}
