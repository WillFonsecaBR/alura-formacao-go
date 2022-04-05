package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 5
const delay = 5

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
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) // Esta função proporciona uma resposta de sucesso para o SO
		default:
			fmt.Println("OPÇÃO INVALIDA")
			os.Exit(-1) // Caso haja algum erro a função retorna este erro e sai
		}
	}
}
func lerSiatesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		os.Exit(-1)
	}

	// Utilizando o Bufio para fazer a leitura do conteúdo.
	leitor := bufio.NewReader(arquivo)
	for {
		// Vai ser lido até a quebra de linha
		linha, err := leitor.ReadString('\n')

		// Removendo os espaços e as quebras de linha do arquivo
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	// É boa pratica fechar o arquivo
	arquivo.Close()
	return sites
}

// Executa o monitoramento
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site: ", site, "Foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("O Site: ", site, "Esta com problemas. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

// Inicia o monitoramento
func iniciarMonitoramento() {
	fmt.Println()
	fmt.Println("=============================")
	fmt.Println()

	fmt.Println("Monitorando aplicação...")

	sites := lerSiatesDoArquivo()

	for i := 0; i < monitoramento; i++ { // Controlar a quantidade de monitoramentos.
		for i, site := range sites {
			fmt.Println("Testando site: ", i, ":", site)
			fmt.Println()
			testaSite(site) // Chama a função testa site.
			fmt.Println()
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println()
	fmt.Println("=============================")
	fmt.Println()
}

// Exibe a introdução do programa
func exibeIntroducao() {
	nome := "Willian Alves Fonseca"
	versao := 1.0

	fmt.Println("=====| INTRODUÇÃO |=====")
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa esta na versão: ", versao)
	fmt.Println("========================")
	fmt.Println()
}

// Exibe as opções de navegação
func exibeMenu() {
	fmt.Println("========| MENU |========")
	fmt.Println("1 - Iníciando do monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")
	fmt.Println("========================")
	fmt.Print("Digite sua opção: ")
}

// Faz a leitura do comando
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func registraLog(site string, status bool) {
	// Estes parametros abre o arquivo, cria um arquivo e adiciona elementos no arquivo.
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}
