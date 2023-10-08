package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(50)
	PagarBoleto(&contaDaLuisa, 45)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

/*
func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.222.33",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100,
	}

	fmt.Println(contaDoBruno)

}
*/
/*
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
*/

//contaDaDani.titular = "Dani"
//contaDaDani := ContaCorrente{}
//contaDaDani.saldo = 50Println(a)

//fmt.Println(contaDaDani.saldo)
//status, valor := contaDaDani.Depositar(300)
//fmt.Println(status, valor)

/*
	fmt.Println(contaDoRoberto)
	fmt.Println(contaDaRaquel)

	var contaDaDaniela *ContaCorrente
	contaDaDaniela = new(ContaCorrente)
	contaDaDaniela.titular = "Daniela"
	contaDaDaniela.saldo = 500
	fmt.Println(*contaDaDaniela)
*/
