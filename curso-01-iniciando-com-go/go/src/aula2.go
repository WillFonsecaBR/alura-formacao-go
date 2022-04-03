package main

import "fmt"

func main() {
	nome := "Willian Alves Fonseca"
	versao := 1.0

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa esta na versão: ", versao)

	// Menu de opções:
	fmt.Println("1 - Iníkco do monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("3 - Sair do programa.")

	// Recebendo imput de dados
	var comando int
	fmt.Scan(&comando)
}
