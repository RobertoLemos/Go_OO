package contas

import "banco/contas/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insulficiente"
	}

}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "alor do deposito é menor que zero", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTranferencia < c.saldo && valorDaTranferencia > 0 {
		c.saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
