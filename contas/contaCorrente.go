package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insulficiente"
	}

}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.Saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "alor do deposito é menor que zero", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTranferencia < c.Saldo && valorDaTranferencia > 0 {
		c.Saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true
	} else {
		return false
	}
}
