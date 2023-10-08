package contas

import (
	"banco/contas/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insulficiente"
	}

}

func (c *ContaPoupanca) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "alor do deposito é menor que zero", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
