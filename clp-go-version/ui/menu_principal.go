package ui

import (
	"bufio"
	"fmt"
	"strconv"
)

// MenuPrincipal representa o menu principal do sistema.
// Ele permite que o usuário escolha entre gerenciar produtos, gerenciar vendas ou encerrar o programa.
type MenuPrincipal struct {
	MenuProduto *MenuProduto // Menu para gerenciar produtos
	MenuVenda   *MenuVenda   // Menu para gerenciar vendas
}

// NewMenuPrincipal cria uma nova instância de MenuPrincipal.
// Retorna um ponteiro para uma instância do menu principal com os menus de produto e venda configurados.
func NewMenuPrincipal() *MenuPrincipal {
	return &MenuPrincipal{
		MenuProduto: NewMenuProduto(),
		MenuVenda:   NewMenuVenda(),
	}
}

// MostrarTitulo exibe o título do menu principal.
// Exibe "MENU PRINCIPAL" na tela.
func (m *MenuPrincipal) MostrarTitulo() {
	fmt.Println("MENU PRINCIPAL")
}

// MostrarOpcoes exibe as opções disponíveis no menu principal.
// Exibe as opções de: fechar o programa, gerenciar produtos ou gerenciar vendas.
func (m *MenuPrincipal) MostrarOpcoes() {
	fmt.Println("0 -> FECHAR PROGRAMA")
	fmt.Println("1 -> PRODUTO")
	fmt.Println("2 -> VENDA")
}

// ExecutarOpcao executa a ação correspondente à opção escolhida pelo usuário.
// Dependendo da opção, o menu de produtos ou vendas é exibido, ou o programa é encerrado.
func (m *MenuPrincipal) ExecutarOpcao(opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		// Caso 0: Exibe mensagem e encerra o programa retornando 0.
		fmt.Println("Programa encerrado.")
		return 0
	case 1:
		// Caso 1: Exibe o menu de produtos chamando o método MostrarMenu do MenuProduto.
		m.MenuProduto.MostrarMenu(scanner)
	case 2:
		// Caso 2: Exibe o menu de vendas chamando o método MostrarMenu do MenuVenda.
		m.MenuVenda.MostrarMenu(scanner)
	default:
		// Caso de opção inválida: Exibe uma mensagem de erro.
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	// Retorna 1 para continuar exibindo o menu.
	return 1
}

// MostrarMenu exibe o menu principal e gerencia as escolhas do usuário.
// Este método exibe o título e as opções, processa a entrada do usuário e executa a opção escolhida.
func (m *MenuPrincipal) MostrarMenu(scanner *bufio.Scanner) {
	for {
		// Exibe o título e as opções do menu principal.
		m.MostrarTitulo()
		m.MostrarOpcoes()

		// Solicita a entrada do usuário.
		fmt.Print("INFORME A SUA OPÇÃO: ")
		if scanner.Scan() {
			// Converte a entrada para um número inteiro e executa a opção.
			opcao, err := strconv.Atoi(scanner.Text())
			if err == nil && m.ExecutarOpcao(opcao, scanner) == 0 {
				// Caso a opção seja 0 (encerrar), o loop é interrompido.
				break
			}
		}
	}
}
