package main

import (
	"fmt"
	"learning-go/Alura/curso-2-poo/banco/clientes"
	"learning-go/Alura/curso-2-poo/banco/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valorDoSaque float64) bool
}

func main() {

	contaBruno := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bruno",
			CPF:       "123.456.789-00",
			Profissao: "Engenheiro de Software",
		},
		NumeroAgencia: 123,
		NumeroConta:   456,
	}

	contaLucas := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Lucas",
			CPF:       "987.654.321-00",
			Profissao: "Engenheiro de Software",
		},
		NumeroAgencia: 123,
		NumeroConta:   456,
	}

	PagarBoleto(&contaBruno, 100)	

	fmt.Println(contaBruno)

	contaBruno.Transferir(100.0, &contaLucas)

	contaBruno.Depositar(100.0)
	contaBruno.Sacar(300)
	fmt.Println(contaBruno)

}
