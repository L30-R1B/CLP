package ui

import (
	"bufio"
	"clp-go-version/data"
	"clp-go-version/entidades"
	"fmt"
	"strconv"
)

// MenuProduto representa o menu para gerenciamento de produtos.
type MenuProduto struct {
	dao *data.DAOProduto
}

// NewMenuProduto cria uma nova instância de MenuProduto.
func NewMenuProduto() *MenuProduto {
	return &MenuProduto{
		dao: data.GetInstance(),
	}
}

// MostrarTitulo exibe o título do menu de produtos.
func (m *MenuProduto) MostrarTitulo() {
	fmt.Println("MENU PRODUTOS")
}

// MostrarOpcoes exibe as opções disponíveis no menu.
func (m *MenuProduto) MostrarOpcoes() {
	fmt.Println("0 -> VOLTAR")
	fmt.Println("1 -> LISTAR")
	fmt.Println("2 -> ADICIONAR")
	fmt.Println("3 -> REMOVER")
}

// MostrarMenu exibe o menu e gerencia as opções.
func (m *MenuProduto) MostrarMenu(scanner *bufio.Scanner) {
	for {
		m.MostrarTitulo()
		m.MostrarOpcoes()

		fmt.Print("INFORME A SUA OPCAO: ")
		scanner.Scan()
		opcao, _ := strconv.Atoi(scanner.Text())

		if m.ExecutarOpcao(opcao, scanner) == 0 {
			break
		}
	}
}

// ExecutarOpcao executa a opção escolhida pelo usuário.
func (m *MenuProduto) ExecutarOpcao(opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		return 0
	case 1:
		m.Listar()
	case 2:
		m.Adicionar(scanner)
	case 3:
		m.Remover(scanner)
	default:
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	return 1
}

// Listar exibe todos os produtos cadastrados no sistema.
func (m *MenuProduto) Listar() {
	fmt.Println(m.dao.String())
}

// Adicionar adiciona um novo produto ao sistema.
func (m *MenuProduto) Adicionar(scanner *bufio.Scanner) {
	var nome string
	var valor float64

	for {
		fmt.Print("\nDigite o nome: ")
		scanner.Scan()
		nome = scanner.Text()

		fmt.Print("Digite o valor: ")
		scanner.Scan()
		valor, _ = strconv.ParseFloat(scanner.Text(), 64)

		if nome == "" || valor <= 0.0 {
			fmt.Println("\nFavor informar os dados corretamente.\n")
			continue
		}
		break
	}

	produto := entidades.NewProduto(nome, valor)
	m.dao.Adicionar(produto)
	fmt.Println("Produto adicionado com sucesso!")
}

// Remover remove um produto com base no nome.
func (m *MenuProduto) Remover(scanner *bufio.Scanner) {
	var nome string

	for {
		fmt.Print("\nDigite o nome: ")
		scanner.Scan()
		nome = scanner.Text()

		if nome == "" {
			fmt.Println("\nFavor informar o nome corretamente.\n")
			continue
		}
		break
	}

	m.dao.Remover(m.dao.BuscarPorNome(nome).GetID())
}
