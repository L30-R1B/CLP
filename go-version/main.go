package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bem-vindo ao sistema de gerenciamento!")

	menu := NewMainMenu()
	menu.ShowMenu()

	fmt.Println("Programa encerrado.")
	os.Exit(0)
}
