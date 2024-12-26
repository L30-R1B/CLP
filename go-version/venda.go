package main

import (
	"fmt"
	"time"
)

type ItemVenda struct {
	Produto    Produto
	Quantidade int
	Valor      float64
}

func NewItemVenda(produto Produto, quantidade int) ItemVenda {
	return ItemVenda{
		Produto:    produto,
		Quantidade: quantidade,
		Valor:      produto.Valor,
	}
}

func (i ItemVenda) String() string {
	return fmt.Sprintf("%15s %8.2f x %5d = %8.2f", i.Produto.Nome, i.Valor, i.Quantidade, i.Valor*float64(i.Quantidade))
}

type Venda struct {
	Entidade
	DataHora time.Time
	Itens    []ItemVenda
}

func NewVenda() Venda {
	return Venda{
		Entidade: NewEntidade(),
		DataHora: time.Now(),
		Itens:    []ItemVenda{},
	}
}

func (v *Venda) AdicionarItem(produto Produto, qtd int) {
	v.Itens = append(v.Itens, NewItemVenda(produto, qtd))
}

func (v *Venda) RemoverItem(nomeProduto string) {
	for index, item := range v.Itens {
		if item.Produto.Nome == nomeProduto {
			v.Itens = append(v.Itens[:index], v.Itens[index+1:]...)
			break
		}
	}
}

func (v Venda) Total() float64 {
	total := 0.0
	for _, item := range v.Itens {
		total += item.Valor * float64(item.Quantidade)
	}
	return total
}

func (v Venda) String() string {
	result := fmt.Sprintf("%sData-Hora: %s\nItens:\n", v.Entidade.String(), v.DataHora.Format(time.RFC1123))
	for _, item := range v.Itens {
		result += "\n  " + item.String()
	}
	result += fmt.Sprintf("\nTOTAL: %.2f", v.Total())
	return result
}
