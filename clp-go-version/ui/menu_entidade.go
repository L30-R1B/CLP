package ui

import (
	"bufio"
	"fmt"
)

// MenuEntidade define um menu abstrato para gerenciar entidades.
// Ele estende a interface MenuAbstrato e adiciona métodos específicos para listar, adicionar e remover entidades.
type MenuEntidade interface {
	MenuAbstrato                      // Métodos da interface MenuAbstrato, como MostrarMenu e MostrarTitulo
	Listar()                          // Método para listar as entidades
	Adicionar(scanner *bufio.Scanner) // Método para adicionar uma nova entidade
	Remover(scanner *bufio.Scanner)   // Método para remover uma entidade
}

// ExecutarOpcao executa a opção escolhida pelo usuário.
// Dependendo da opção, ele chama os métodos apropriados do menu.
func ExecutarOpcao(menu MenuEntidade, opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		// Caso 0: Retorna 0 para indicar que o menu deve ser fechado.
		return 0
	case 1:
		// Caso 1: Chama o método Listar para exibir as entidades.
		menu.Listar()
	case 2:
		// Caso 2: Chama o método Adicionar para adicionar uma nova entidade.
		menu.Adicionar(scanner)
	case 3:
		// Caso 3: Chama o método Remover para remover uma entidade.
		menu.Remover(scanner)
	default:
		// Caso inválido: Informa que a opção escolhida é inválida.
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	// Retorna 1 para continuar exibindo o menu.
	return 1
}

// MostrarOpcoes exibe as opções disponíveis no menu.
// Este método exibe as opções que o usuário pode escolher.
func MostrarOpcoes() {
	// Opções do menu: voltar, listar, adicionar e remover.
	fmt.Println("0 -> VOLTAR")
	fmt.Println("1 -> LISTAR")
	fmt.Println("2 -> ADICIONAR")
	fmt.Println("3 -> REMOVER")
}
