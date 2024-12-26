package main

import "fmt"

type Produto struct {
	Entidade
	Nome  string
	Valor float64
}

func NewProduto(nome string, valor float64) Produto {
	return Produto{
		Entidade: NewEntidade(),
		Nome:     nome,
		Valor:    valor,
	}
}

func (p Produto) String() string {
	return fmt.Sprintf("%sNome: %s\tValor: %.2f", p.Entidade.String(), p.Nome, p.Valor)
}
