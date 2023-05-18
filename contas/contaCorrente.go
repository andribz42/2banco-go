package contas

import "estudo/banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, conta *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.Sacar(valorDaTransferencia)
		conta.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0 && c.saldo > 0

	if valorDoSaque < 0 {
		return "Valor do saque é inválido."
	}

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado."
	} else {
		return "Sem saldo para saque."
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito feito com sucesso!"
	}
	return "Valor do depósito inválido."
}
