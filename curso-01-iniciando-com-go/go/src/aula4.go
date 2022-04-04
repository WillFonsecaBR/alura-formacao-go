package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for { //Este tipo de for roda infinitamente pois não tem While em go.
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
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
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando aplicação...")

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O Site: ", site, "Foi carregado com sucesso")
	} else {
		fmt.Println("O Site: ", site, "Esta com problemas. Status Code: ", resp.StatusCode)
	}

}

func exibeIntroducao() {
	nome := "Willian Alves Fonseca"
	versao := 1.0

	fmt.Println("=====| INTRODUÇÃO |=====")
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa esta na versão: ", versao)
	fmt.Println("========================")
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("========| MENU |========")
	fmt.Println("1 - Iníciando do monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")
	fmt.Println("========================")
	fmt.Print("Digite sua opção: ")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
