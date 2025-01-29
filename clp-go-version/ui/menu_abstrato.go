package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MenuAbstrato define uma interface para menus abstratos.
type MenuAbstrato interface {
	MostrarMenu()                                        // Exibe o menu.
	ExecutarOpcao(opcao int, scanner *bufio.Scanner) int // Executa a opção escolhida.
	MostrarOpcoes()                                      // Exibe as opções do menu.
	MostrarTitulo()                                      // Exibe o título do menu.
}

// MostrarMenu exibe o menu e gerencia as opções escolhidas pelo usuário.
func MostrarMenu(menu MenuAbstrato) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n\n\n")
		menu.MostrarTitulo()
		menu.MostrarOpcoes()

		fmt.Print("INFORME A SUA OPÇÃO: ")
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			opcao, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Entrada inválida. Por favor, insira um número.")
				continue
			}

			if menu.ExecutarOpcao(opcao, scanner) == 0 {
				break
			}
		}
	}
}
