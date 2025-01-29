package entidades

import (
	"fmt"
	"strings"
	"time"
)

// Venda representa uma venda com data, hora e itens.
type Venda struct {
	ID       int64
	DataHora time.Time
	Itens    []ItemVenda
}

// NewVenda cria uma nova instância de Venda.
func NewVenda() *Venda {
	return &Venda{
		ID:       time.Now().UnixMilli(),
		DataHora: time.Now(),
		Itens:    []ItemVenda{},
	}
}

// GetID retorna o ID da Venda.
func (v *Venda) GetID() int64 {
	return v.ID
}

// String retorna uma representação textual da Venda.
func (v *Venda) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Venda[ID=%d, DataHora=%s]\n", v.ID, v.DataHora.Format("2006-01-02 15:04:05")))
	sb.WriteString("Itens:\n")
	for _, item := range v.Itens {
		sb.WriteString(fmt.Sprintf("  %s\n", item.String()))
	}
	sb.WriteString(fmt.Sprintf("TOTAL: %.2f\n", v.Total()))
	return sb.String()
}

// GetDataHora retorna a data e hora da Venda.
func (v *Venda) GetDataHora() time.Time {
	return v.DataHora
}

// GetItens retorna a lista de itens da Venda.
func (v *Venda) GetItens() []ItemVenda {
	return v.Itens
}

// AdicionarItem adiciona um novo item à Venda.
func (v *Venda) AdicionarItem(produto Produto, quantidade int) {
	v.Itens = append(v.Itens, ItemVenda{
		Produto:    produto,
		Quantidade: quantidade,
		Valor:      produto.GetValor(),
	})
}

// RemoverItemPorPosicao remove um item da Venda com base na sua posição na lista.
func (v *Venda) RemoverItemPorPosicao(posicao int) {
	if posicao >= 0 && posicao < len(v.Itens) {
		v.Itens = append(v.Itens[:posicao], v.Itens[posicao+1:]...)
	}
}

// RemoverItemPorNome remove um item da Venda com base no nome do produto.
func (v *Venda) RemoverItemPorNome(nomeProduto string) {
	filtrados := []ItemVenda{}
	for _, item := range v.Itens {
		if strings.EqualFold(item.Produto.GetNome(), nomeProduto) {
			continue
		}
		filtrados = append(filtrados, item)
	}
	v.Itens = filtrados
}

// Total calcula o valor total da Venda.
func (v *Venda) Total() float64 {
	total := 0.0
	for _, item := range v.Itens {
		total += item.Valor * float64(item.Quantidade)
	}
	return total
}
