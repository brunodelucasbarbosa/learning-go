package main

import "fmt"

type pessoa struct {
	nome      string
	idade     int
	profissao string
	endereco  endereco
}

type endereco struct {
	logradouro string
	numero     int
	cidade     string
	estado     string
	cep        string
}

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	item1 := item{produtoID: 1, quantidade: 2, preco: 10.0}
	item2 := item{produtoID: 2, quantidade: 3, preco: 20.0}
	item3 := item{produtoID: 3, quantidade: 4, preco: 30.0}
	item4 := item{produtoID: 4, quantidade: 5, preco: 40.0}

	pedido1 := pedido{
		userID: 1,
		itens: []item{
			item1,
			item2,
			item3,
			item4,
		},
	}

	fmt.Println(pedido1.valorTotal())

	pessoa1 := pessoa{
		nome:      "João",
		idade:     30,
		profissao: "Programador",
		endereco: endereco{
			logradouro: "Rua ABC",
			numero:     123,
			cidade:     "São Paulo",
			estado:     "SP",
			cep:        "12345-678",
		},
	}

	fmt.Println(pessoa1)
}
