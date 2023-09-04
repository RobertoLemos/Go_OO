package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoRoberto := ContaCorrente{
		titular:       "Roberto",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}
	contaDaRaquel := ContaCorrente{
		"Raquel",
		222,
		111222,
		200,
	}
	fmt.Println(contaDoRoberto)
	fmt.Println(contaDaRaquel)

}
