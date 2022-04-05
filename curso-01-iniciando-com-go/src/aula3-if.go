package main

import "fmt"

func main() {
	nome := "Willian Alves Fonseca"
	versao := 1.0

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa esta na versão: ", versao)

	// MENU DE OPÇÕES:
	fmt.Println("1 - Iníciando do monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")

	// RECEBENDO IMPUT DE DADOS:
	var comando int
	fmt.Scan(&comando)
	fmt.Println("==========================")
	// &comando Mostra o caminho de alocação na memoria

	// CRIANDO FLUXO DE VALIDAÇÃO DE ESCOLHA:
	if comando == 1 {
		fmt.Println("Monitorando aplicação...")
	} else if comando == 2 {
		fmt.Println("Exibindo LOGs da aplicação...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("OPÇÃO INVALIDA")
	}
}
