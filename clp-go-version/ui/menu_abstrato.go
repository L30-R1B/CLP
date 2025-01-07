package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MenuAbstrato define uma interface para menus abstratos.
// Essa interface permite implementar menus customizados com comportamento específico.
type MenuAbstrato interface {
	MostrarMenu()                                        // Método para exibir o menu.
	ExecutarOpcao(opcao int, scanner *bufio.Scanner) int // Método para executar uma opção.
	MostrarOpcoes()                                      // Método para exibir as opções do menu.
	MostrarTitulo()                                      // Método para exibir o título do menu.
}

// MostrarMenu exibe o menu e gerencia as opções escolhidas pelo usuário.
// Ele utiliza a interface MenuAbstrato para trabalhar com menus de forma genérica.
func MostrarMenu(menu MenuAbstrato) {
	scanner := bufio.NewScanner(os.Stdin) // Cria um scanner para leitura de entrada do usuário.
	for {
		// Exibe espaçamento para clareza na saída do terminal.
		fmt.Println("\n\n\n")
		menu.MostrarTitulo() // Exibe o título do menu.
		menu.MostrarOpcoes() // Exibe as opções disponíveis.

		// Solicita ao usuário que informe sua escolha.
		fmt.Print("INFORME A SUA OPÇÃO: ")
		if scanner.Scan() { // Lê a entrada do usuário.
			input := strings.TrimSpace(scanner.Text()) // Remove espaços em branco.
			opcao, err := strconv.Atoi(input)          // Converte a entrada para um inteiro.
			if err != nil {
				fmt.Println("Entrada inválida. Por favor, insira um número.")
				continue // Retorna ao início do loop em caso de erro de conversão.
			}

			// Executa a opção escolhida e verifica o retorno.
			if menu.ExecutarOpcao(opcao, scanner) == 0 {
				break // Sai do loop se o retorno for 0.
			}
		}
	}
}
