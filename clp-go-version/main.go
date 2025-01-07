package main

import (
	"bufio"
	"clp-go-version/ui"
	"fmt"
	"os"
)

func main() {
	// Cria uma instância do scanner para leitura do input do usuário.
	scanner := bufio.NewScanner(os.Stdin)

	// Cria uma instância de MenuPrincipal e chama o método MostrarMenu.
	menuPrincipal := ui.NewMenuPrincipal()
	menuPrincipal.MostrarMenu(scanner)

	fmt.Println("Programa encerrado.")
}
