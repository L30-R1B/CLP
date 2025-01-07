package data

import (
	"clp-go-version/entidades"
	"sync"
)

// DAOProduto é um singleton para gerenciar o DAO de Produto.
// Ele encapsula o DAO genérico especializado para produtos e garante uma instância única.
type DAOProduto struct {
	dao *DAO[*entidades.Produto] // DAO genérico para a entidade Produto.
}

var instance *DAOProduto // Instância única do singleton DAOProduto.
var once sync.Once       // Garantia de inicialização única e thread-safe.

// GetInstance retorna a instância singleton de DAOProduto.
// A instância é inicializada apenas uma vez usando o sync.Once.
func GetInstance() *DAOProduto {
	once.Do(func() {
		instance = &DAOProduto{
			dao: NewDAO[*entidades.Produto](), // Criação do DAO especializado para Produto.
		}
	})
	return instance
}

// Adicionar adiciona um Produto ao DAO.
// Este método encapsula a lógica de adição diretamente no DAO genérico.
func (d *DAOProduto) Adicionar(produto *entidades.Produto) {
	d.dao.Adicionar(produto)
}

// Buscar por ID retorna um Produto com o ID especificado.
// Realiza a busca no DAO genérico e retorna o ponteiro do produto correspondente.
func (d *DAOProduto) Buscar(id int64) *entidades.Produto {
	return *d.dao.Buscar(id) // Retorna o ponteiro desreferenciado do produto.
}

// BuscarPorNome retorna um Produto com o nome especificado.
// Realiza a busca iterando sobre os dados armazenados no DAO.
func (d *DAOProduto) BuscarPorNome(nome string) *entidades.Produto {
	for _, p := range d.dao.GetDados() {
		if p.GetNome() == nome { // Compara o nome do produto.
			return p
		}
	}
	return nil // Retorna nil caso não encontre.
}

// Remover por ID remove um Produto com o ID especificado.
// Encapsula a lógica de remoção no DAO genérico.
func (d *DAOProduto) Remover(id int64) {
	d.dao.Remover(id)
}

// RemoverPorNome remove um Produto com o nome especificado.
// Filtra os produtos e mantém apenas aqueles cujo nome não corresponde ao fornecido.
func (d *DAOProduto) RemoverPorNome(nome string) {
	filtrados := []*entidades.Produto{}
	for _, p := range d.dao.GetDados() {
		if p.GetNome() != nome { // Compara o nome para exclusão.
			filtrados = append(filtrados, p)
		}
	}
	d.dao.Dados = filtrados // Atualiza a lista de dados no DAO.
}

// String retorna uma representação textual do DAO de Produtos.
// Utiliza o método String do DAO genérico para formatar os produtos armazenados.
func (d *DAOProduto) String() string {
	return d.dao.String()
}
