package main

import "fmt"

// O struct é equivalente a
type contaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

// Função para sacar
func (c *contaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saldo realizado com sucesso!"
	} else {
		return "saldo insuficiente"
	}
}

func main() {
	// Forma mais usual de utilizar o struct
	contaGuilherme := contaCorrente{"Willian Fonseca", 296, 32, 4988.30}
	fmt.Println(contaGuilherme)

	// Outra forma de instanciar e usar o instruct.
	// Nesta forma o * é usado como direcionamento, um ponteiro para a referencia da variavel.
	var contaDaCris *contaCorrente
	contaDaCris = new(contaCorrente)
	contaDaCris.titular = "Cris Oliveira"
	contaDaCris.agencia = 326
	contaDaCris.conta = 3265045
	contaDaCris.saldo = 6000.32
	fmt.Println(contaDaCris)

	fmt.Println(contaDaCris.Sacar(1200.00))
	fmt.Println(contaDaCris)
}
