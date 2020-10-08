package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p *pessoa) falar () {
	fmt.Println("Me chamo", p.nome, p.sobrenome, "tenho", p.idade, "anos")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	eu := pessoa{"Danilo", "Macedo", 25}

	var p *pessoa
	p = &eu

	dizerAlgumaCoisa(p)
	// dizerAlgumaCoisa(eu)
}