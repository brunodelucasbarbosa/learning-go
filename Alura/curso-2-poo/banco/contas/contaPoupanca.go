package contas

import "learning-go/Alura/curso-2-poo/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) bool {
	if valorDoSaque <= c.saldo && valorDoSaque > 0 {
		c.saldo -= valorDoSaque
		return true
	}
	return false
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	c.saldo += valorDoDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)
		return true
	}
	return false
}

func (c *ContaPoupanca) getSaldo() float64 {
	return c.saldo
}
