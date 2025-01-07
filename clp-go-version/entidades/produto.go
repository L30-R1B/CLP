package entidades

import (
	"fmt"  // O pacote `fmt` é usado para formatação e saída de strings.
	"time" // O pacote `time` fornece funcionalidades para trabalhar com data e hora.
)

// Produto representa um produto com nome e valor.
// Em Go, structs são usadas para agrupar campos relacionados. São semelhantes a classes em outras linguagens,
// mas Go não possui herança. Em vez disso, utiliza composição para reutilização de código.
type Produto struct {
	ID    int64   // Campo para armazenar o identificador único do produto. Aqui utilizamos `int64` para garantir precisão.
	Nome  string  // Nome do produto, armazenado como uma string.
	Valor float64 // Valor do produto, representado com precisão decimal usando o tipo `float64`.
}

// ItemVenda representa um item em uma venda.
// Essa struct usa composição para relacionar um produto a uma venda, incluindo a quantidade e o valor total.
type ItemVenda struct {
	Produto    Produto // Um campo que referencia a struct Produto.
	Quantidade int     // Número de unidades do produto na venda.
	Valor      float64 // Valor total desse item, calculado como `Quantidade * Produto.Valor`.
}

// NewProduto cria um novo Produto com valores padrão.
// Funções iniciadas com "New" são convenções em Go para criar e inicializar structs.
// Essa função retorna um ponteiro para um novo Produto.
func NewProduto(nome string, valor float64) *Produto {
	return &Produto{
		ID:    time.Now().UnixMilli(), // Gera um ID único usando o timestamp em milissegundos.
		Nome:  nome,                   // Inicializa o campo Nome com o valor fornecido.
		Valor: valor,                  // Inicializa o campo Valor com o valor fornecido.
	}
}

// GetID retorna o ID do Produto.
// Métodos em Go são declarados ao associá-los a um tipo específico usando um receptor.
// Aqui usamos um receptor `*Produto` (ponteiro), o que é eficiente para structs grandes,
// além de permitir modificar o objeto se necessário.
func (p *Produto) GetID() int64 {
	return p.ID // Retorna o identificador único do produto.
}

// String retorna uma representação textual do Produto.
// Esse método implementa a interface `fmt.Stringer`, o que permite formatar um Produto em strings personalizadas.
func (p *Produto) String() string {
	return fmt.Sprintf("Produto[ID=%d, Nome=%s, Valor=%.2f]", p.ID, p.Nome, p.Valor)
}

// GetNome retorna o nome do Produto.
// Métodos como este são úteis para encapsular a lógica de acesso, mesmo que seja simples.
func (p *Produto) GetNome() string {
	return p.Nome
}

// GetValor retorna o valor do Produto.
// Em Go, a abordagem de encapsulamento é opcional, mas pode ser usada para garantir consistência ou validação.
func (p *Produto) GetValor() float64 {
	return p.Valor
}

// SetNome define o nome do Produto.
// Métodos `Set` permitem modificar campos de uma struct de forma controlada.
// Aqui estamos aceitando diretamente o valor sem validação.
func (p *Produto) SetNome(nome string) {
	p.Nome = nome // Atualiza o campo Nome com o valor fornecido.
}

// SetValor define o valor do Produto.
// Este método é semelhante ao `SetNome`, mas modifica o campo Valor.
// Ele poderia ser estendido para incluir validações, como verificar se o valor não é negativo.
func (p *Produto) SetValor(valor float64) {
	p.Valor = valor // Atualiza o campo Valor com o valor fornecido.
}
