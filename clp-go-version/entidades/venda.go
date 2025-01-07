package entidades

import (
	"fmt"     // O pacote `fmt` é usado para formatação e saída de strings.
	"strings" // O pacote `strings` oferece utilitários para manipulação de strings.
	"time"    // O pacote `time` fornece funcionalidades para lidar com datas e horários.
)

// Venda representa uma venda com data, hora e itens.
// Essa struct encapsula informações sobre uma transação, incluindo seus itens e o momento em que foi realizada.
type Venda struct {
	ID       int64       // Identificador único da venda, gerado com base no timestamp.
	DataHora time.Time   // Data e hora da venda, armazenada no formato `time.Time`.
	Itens    []ItemVenda // Lista de itens associados à venda, usando a struct `ItemVenda`.
}

// NewVenda cria uma nova instância de Venda.
// A função utiliza o timestamp atual para gerar um ID único e inicializa a lista de itens como vazia.
func NewVenda() *Venda {
	return &Venda{
		ID:       time.Now().UnixMilli(), // Gera um ID único com base no tempo atual em milissegundos.
		DataHora: time.Now(),             // Define a data e hora atual no campo `DataHora`.
		Itens:    []ItemVenda{},          // Inicializa a lista de itens como vazia.
	}
}

// GetID retorna o ID da Venda.
// Este método encapsula o acesso ao identificador único da venda.
func (v *Venda) GetID() int64 {
	return v.ID // Retorna o valor do campo `ID`.
}

// String retorna uma representação textual da Venda.
// Implementa a interface `fmt.Stringer` para personalizar a saída textual da venda.
func (v *Venda) String() string {
	return fmt.Sprintf("Venda[ID=%d, DataHora=%s]",
		v.ID,
		v.DataHora.Format("2006-01-02 15:04:05"), // Formata a data e hora no padrão `YYYY-MM-DD HH:mm:ss`.
	)
}

// GetDataHora retorna a data e hora da Venda.
// Encapsula o acesso ao campo `DataHora`, que contém o momento em que a venda foi criada.
func (v *Venda) GetDataHora() time.Time {
	return v.DataHora // Retorna o valor do campo `DataHora`.
}

// GetItens retorna a lista de itens da Venda.
// Essa função fornece acesso à lista de itens que compõem a venda.
func (v *Venda) GetItens() []ItemVenda {
	return v.Itens // Retorna a fatia (slice) de itens.
}

// AdicionarItem adiciona um novo item à Venda.
// Este método cria um novo `ItemVenda` a partir de um produto e uma quantidade especificada e o adiciona à lista de itens.
func (v *Venda) AdicionarItem(produto Produto, quantidade int) {
	v.Itens = append(v.Itens, ItemVenda{
		Produto:    produto,            // Produto associado ao item.
		Quantidade: quantidade,         // Quantidade especificada.
		Valor:      produto.GetValor(), // Valor unitário do produto.
	})
}

// RemoverItemPorPosicao remove um item da Venda com base na sua posição na lista.
// Utiliza a função `append` para criar uma nova lista excluindo o item na posição especificada.
func (v *Venda) RemoverItemPorPosicao(posicao int) {
	if posicao >= 0 && posicao < len(v.Itens) { // Verifica se a posição é válida.
		v.Itens = append(v.Itens[:posicao], v.Itens[posicao+1:]...) // Remove o item na posição.
	}
}

// RemoverItemPorNome remove um item da Venda com base no nome do produto.
// A busca pelo nome é case-insensitive graças à função `strings.EqualFold`.
func (v *Venda) RemoverItemPorNome(nomeProduto string) {
	filtrados := []ItemVenda{} // Inicializa uma nova lista para itens filtrados.
	for _, item := range v.Itens {
		if strings.EqualFold(item.Produto.GetNome(), nomeProduto) {
			continue // Ignora o item se o nome do produto corresponder ao nome fornecido.
		}
		filtrados = append(filtrados, item) // Adiciona o item à nova lista se não houver correspondência.
	}
	v.Itens = filtrados // Atualiza a lista de itens com os filtrados.
}

// Total calcula o valor total da Venda.
// O valor total é calculado multiplicando o valor unitário pela quantidade de cada item.
func (v *Venda) Total() float64 {
	total := 0.0 // Inicializa o total como 0.
	for _, item := range v.Itens {
		total += item.Valor * float64(item.Quantidade) // Soma o valor total de cada item.
	}
	return total // Retorna o valor total da venda.
}
