package main

import (
	"fmt"

	"banco/contas"
)

func main() {
	contaDaDani := contas.ContaCorrente{Titular: "Dani", Saldo: 300}
	contaDaAntonella := contas.ContaCorrente{Titular: "Antonella", Saldo: 100}

	status := contaDaDani.Transferir(200, &contaDaAntonella)

	fmt.Println(status)
	fmt.Println(contaDaDani)
	fmt.Println(contaDaAntonella)
}

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
