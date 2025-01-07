package data

import (
	"clp-go-version/entidades"
	"sync"
)

// DAOVenda é um singleton para gerenciar o DAO de Venda.
// Encapsula o DAO específico para a entidade Venda e garante que apenas uma instância seja criada.
type DAOVenda struct {
	dao *DAO[*entidades.Venda] // Referência ao DAO genérico, especializado para vendas.
}

var vendaInstance *DAOVenda // Instância única do DAOVenda.
var vendaOnce sync.Once     // Usado para garantir inicialização única e thread-safe do singleton.

// GetVendaInstance retorna a instância singleton de DAOVenda.
// Usa `sync.Once` para garantir que a inicialização ocorra apenas uma vez, mesmo em ambientes concorrentes.
func GetVendaInstance() *DAOVenda {
	vendaOnce.Do(func() {
		// Inicializa a instância única do DAOVenda.
		vendaInstance = &DAOVenda{
			dao: NewDAO[*entidades.Venda](), // Cria um novo DAO especializado para vendas.
		}
	})
	return vendaInstance
}

// Adicionar adiciona uma Venda ao DAO.
// Encapsula a funcionalidade de adicionar uma venda na estrutura de dados do DAO.
func (d *DAOVenda) Adicionar(venda *entidades.Venda) {
	d.dao.Adicionar(venda)
}

// Buscar por ID retorna uma Venda com o ID especificado.
// Realiza a busca no DAO e retorna a referência da venda correspondente.
func (d *DAOVenda) Buscar(id int64) *entidades.Venda {
	return *d.dao.Buscar(id) // Retorna o ponteiro desreferenciado da venda encontrada.
}

// Remover por ID remove uma Venda com o ID especificado.
// Encapsula a funcionalidade de remoção de vendas no DAO.
func (d *DAOVenda) Remover(id int64) {
	d.dao.Remover(id)
}

// String retorna uma representação textual do DAO de Vendas.
// Utiliza o método `String` do DAO para formatar as vendas armazenadas.
func (d *DAOVenda) String() string {
	return d.dao.String()
}
