package ui

import (
	"bufio"
	"fmt"
	"strconv"

	"clp-go-version/data"
	"clp-go-version/entidades"
)

// MenuVenda representa o menu para gerenciamento de vendas.
type MenuVenda struct {
	daoVenda   *data.DAOVenda
	daoProduto *data.DAOProduto
}

// NewMenuVenda cria uma nova instância de MenuVenda.
func NewMenuVenda() *MenuVenda {
	return &MenuVenda{
		daoVenda:   data.GetVendaInstance(),
		daoProduto: data.GetInstance(),
	}
}

// MostrarTitulo exibe o título do menu de vendas.
func (m *MenuVenda) MostrarTitulo() {
	fmt.Println("MENU VENDAS")
}

// MostrarOpcoes exibe as opções disponíveis no menu.
func (m *MenuVenda) MostrarOpcoes() {
	fmt.Println("0 -> VOLTAR")
	fmt.Println("1 -> LISTAR")
	fmt.Println("2 -> ADICIONAR")
	fmt.Println("3 -> REMOVER")
}

// MostrarMenu exibe o menu e gerencia as opções.
func (m *MenuVenda) MostrarMenu(scanner *bufio.Scanner) {
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
func (m *MenuVenda) ExecutarOpcao(opcao int, scanner *bufio.Scanner) int {
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

// Listar exibe todas as vendas cadastradas no sistema.
func (m *MenuVenda) Listar() {
	fmt.Println(m.daoVenda.String())
}

// Adicionar adiciona uma nova venda ao sistema.
func (m *MenuVenda) Adicionar(scanner *bufio.Scanner) {
	venda := entidades.NewVenda()

	for {
		var produto *entidades.Produto
		var qtd int

		for {
			fmt.Print("\nDigite o nome do produto: ")
			scanner.Scan()
			nomeProduto := scanner.Text()

			produto = m.daoProduto.BuscarPorNome(nomeProduto)
			if produto == nil {
				fmt.Println("Produto não encontrado. Tente novamente.")
				continue
			}

			fmt.Print("Digite a quantidade: ")
			scanner.Scan()
			qtd, _ = strconv.Atoi(scanner.Text())

			if qtd <= 0 {
				fmt.Println("Quantidade inválida. Tente novamente.")
				continue
			}
			break
		}

		venda.AdicionarItem(*produto, qtd)

		fmt.Print("\nDeseja adicionar outro produto à venda (1-SIM/0-NAO)? ")
		scanner.Scan()
		opcao, _ := strconv.Atoi(scanner.Text())
		if opcao != 1 {
			break
		}
	}

	fmt.Println("\n\nNOTA FISCAL\n", venda.String())
	m.daoVenda.Adicionar(venda)
}

// Remover remove uma venda com base no ID.
func (m *MenuVenda) Remover(scanner *bufio.Scanner) {
	var id int64

	for {
		fmt.Print("\nDigite o id: ")
		scanner.Scan()
		id, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		if id <= 0 {
			fmt.Println("ID inválido. Tente novamente.")
			continue
		}
		break
	}

	m.daoVenda.Remover(id)
}
