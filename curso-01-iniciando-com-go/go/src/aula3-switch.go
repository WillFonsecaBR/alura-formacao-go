package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	// CRIANDO FLUXO DE VALIDAÇÃO DE ESCOLHA:
	switch comando {
	case 1:
		fmt.Println("Monitorando aplicação...")
		break // Não é necessario osar o BREAK, mas tambem não atrapalha
	case 2:
		fmt.Println("Exibindo LOGs da aplicação...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0) // Esta função proporciona uma resposta de sucesso para o SO
	default:
		fmt.Println("OPÇÃO INVALIDA")
		os.Exit(-1) // Caso haja algum erro a função retorna este erro
	}
}

func exibeIntroducao() {
	nome := "Willian Alves Fonseca"
	versao := 1.0

	fmt.Println("======| INTRODUÇÃO |======")
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa esta na versão: ", versao)
	fmt.Println("==========================")
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("=========| MENU |=========")
	fmt.Println("1 - Iníciando do monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")
	fmt.Println("==========================")
	fmt.Print("Digite sua opção: ")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
