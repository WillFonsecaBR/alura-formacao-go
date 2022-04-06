package main

import "fmt"

// O struct é equivalente a
type contaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {

	// Forma mais usual de utilizar o struct
	contaGuilherme := contaCorrente{"Willian", 296, 32, 4988.30}
	fmt.Println(contaGuilherme)

}
