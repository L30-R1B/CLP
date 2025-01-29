package ui

import (
	"bufio"
	"fmt"
	"strconv"
)

// MenuPrincipal representa o menu principal do sistema.
type MenuPrincipal struct {
	MenuProduto *MenuProduto
	MenuVenda   *MenuVenda
}

// NewMenuPrincipal cria uma nova instância de MenuPrincipal.
func NewMenuPrincipal() *MenuPrincipal {
	return &MenuPrincipal{
		MenuProduto: NewMenuProduto(),
		MenuVenda:   NewMenuVenda(),
	}
}

// MostrarTitulo exibe o título do menu principal.
func (m *MenuPrincipal) MostrarTitulo() {
	fmt.Println("MENU PRINCIPAL")
}

// MostrarOpcoes exibe as opções disponíveis no menu principal.
func (m *MenuPrincipal) MostrarOpcoes() {
	fmt.Println("0 -> FECHAR PROGRAMA")
	fmt.Println("1 -> PRODUTO")
	fmt.Println("2 -> VENDA")
}

// ExecutarOpcao executa a ação correspondente à opção escolhida pelo usuário.
func (m *MenuPrincipal) ExecutarOpcao(opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		return 0
	case 1:
		m.MenuProduto.MostrarMenu(scanner)
	case 2:
		m.MenuVenda.MostrarMenu(scanner)
	default:
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	return 1
}

// MostrarMenu exibe o menu principal e gerencia as escolhas do usuário.
func (m *MenuPrincipal) MostrarMenu(scanner *bufio.Scanner) {
	for {
		m.MostrarTitulo()
		m.MostrarOpcoes()

		fmt.Print("INFORME A SUA OPÇÃO: ")
		if scanner.Scan() {
			opcao, err := strconv.Atoi(scanner.Text())
			if err == nil && m.ExecutarOpcao(opcao, scanner) == 0 {
				break
			}
		}
	}
}
