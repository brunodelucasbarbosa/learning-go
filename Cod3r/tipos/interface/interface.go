package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	idade     int
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return p.nome + " - R$ " + fmt.Sprintf("%.2f", p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var p1 imprimivel = pessoa{"João", 20, "Silva"}
	var p2 imprimivel = produto{"Lapis", 1.80}

	imprimir(p1)
	imprimir(p2)

}
