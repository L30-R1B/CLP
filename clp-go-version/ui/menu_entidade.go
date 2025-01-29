package ui

import (
	"bufio"
	"fmt"
)

// MenuEntidade define um menu para gerenciar entidades.
type MenuEntidade interface {
	MenuAbstrato                      // Herda métodos de MenuAbstrato.
	Listar()                          // Lista as entidades.
	Adicionar(scanner *bufio.Scanner) // Adiciona uma nova entidade.
	Remover(scanner *bufio.Scanner)   // Remove uma entidade.
}

// ExecutarOpcao executa a opção escolhida pelo usuário.
func ExecutarOpcao(menu MenuEntidade, opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		return 0
	case 1:
		menu.Listar()
	case 2:
		menu.Adicionar(scanner)
	case 3:
		menu.Remover(scanner)
	default:
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	return 1
}

// MostrarOpcoes exibe as opções disponíveis no menu.
func MostrarOpcoes() {
	fmt.Println("0 -> VOLTAR")
	fmt.Println("1 -> LISTAR")
	fmt.Println("2 -> ADICIONAR")
	fmt.Println("3 -> REMOVER")
}
