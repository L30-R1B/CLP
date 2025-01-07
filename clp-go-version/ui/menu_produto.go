package ui

import (
	"bufio"
	"clp-go-version/data"
	"clp-go-version/entidades"
	"fmt"
	"strconv"
)

// MenuProduto representa o menu para gerenciamento de produtos.
// Ele oferece opções para listar, adicionar e remover produtos do sistema.
type MenuProduto struct {
	dao *data.DAOProduto // Acesso ao DAO para operações com produtos
}

// NewMenuProduto cria uma nova instância de MenuProduto.
// Retorna um ponteiro para uma instância do menu de produtos com o DAO configurado.
func NewMenuProduto() *MenuProduto {
	return &MenuProduto{
		dao: data.GetInstance(),
	}
}

// MostrarTitulo exibe o título do menu de produtos.
// Este método imprime o título do menu para o usuário.
func (m *MenuProduto) MostrarTitulo() {
	fmt.Println("MENU PRODUTOS")
}

// MostrarOpcoes exibe as opções disponíveis no menu.
// As opções incluem listar, adicionar, remover ou voltar.
func (m *MenuProduto) MostrarOpcoes() {
	fmt.Println("0 -> VOLTAR")
	fmt.Println("1 -> LISTAR")
	fmt.Println("2 -> ADICIONAR")
	fmt.Println("3 -> REMOVER")
}

// MostrarMenu exibe o menu e gerencia as opções.
// Este método exibe o menu e processa as escolhas do usuário em um loop contínuo.
func (m *MenuProduto) MostrarMenu(scanner *bufio.Scanner) {
	var opcao int
	for {
		// Exibe o título e as opções do menu.
		m.MostrarTitulo()
		m.MostrarOpcoes()

		// Solicita a entrada do usuário para escolher uma opção.
		fmt.Print("INFORME A SUA OPCAO: ")
		scanner.Scan()
		opcao, _ = strconv.Atoi(scanner.Text())

		// Executa a opção escolhida pelo usuário.
		if m.ExecutarOpcao(opcao, scanner) == 0 {
			break
		}
	}
}

// ExecutarOpcao executa a opção escolhida pelo usuário.
// Dependendo da opção, chama o método correspondente (listar, adicionar, remover).
func (m *MenuProduto) ExecutarOpcao(opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		// Caso 0: Retorna 0 para sair do menu.
		return 0
	case 1:
		// Caso 1: Chama o método Listar para exibir todos os produtos.
		m.Listar()
	case 2:
		// Caso 2: Chama o método Adicionar para adicionar um novo produto.
		m.Adicionar(scanner)
	case 3:
		// Caso 3: Chama o método Remover para remover um produto.
		m.Remover(scanner)
	default:
		// Caso inválido: Exibe mensagem de erro.
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	// Retorna 1 para continuar exibindo o menu.
	return 1
}

// Listar exibe todos os produtos cadastrados no sistema.
// Este método imprime a lista de produtos do DAO.
func (m *MenuProduto) Listar() {
	// Exibe os dados do DAO de produtos.
	fmt.Println(m.dao.String())
}

// Adicionar adiciona um novo produto ao sistema.
// O método solicita ao usuário o nome e o valor do produto, validando as entradas.
func (m *MenuProduto) Adicionar(scanner *bufio.Scanner) {
	var nome string
	var valor float64

	// Loop para garantir que o nome e valor sejam informados corretamente.
	for {
		// Solicita o nome do produto.
		fmt.Print("\nDigite o nome: ")
		scanner.Scan()
		nome = scanner.Text()

		// Solicita o valor do produto.
		fmt.Print("Digite o valor: ")
		scanner.Scan()
		valor, _ = strconv.ParseFloat(scanner.Text(), 64)

		// Valida se os dados são válidos.
		if nome == "" || valor <= 0.0 {
			// Caso algum dado esteja incorreto, exibe mensagem de erro e solicita novamente.
			fmt.Println("\nFavor informar os dados corretamente.\n")
			continue
		}
		// Sai do loop se os dados estiverem corretos.
		break
	}

	// Cria um novo produto com os dados informados.
	produto := entidades.NewProduto(nome, valor)
	// Adiciona o produto ao DAO.
	m.dao.Adicionar(produto)
	// Exibe mensagem de sucesso.
	fmt.Println("Produto adicionado com sucesso!")
}

// Remover remove um produto com base no nome.
// O método solicita o nome do produto a ser removido e, caso encontrado, o remove do DAO.
func (m *MenuProduto) Remover(scanner *bufio.Scanner) {
	var nome string

	// Loop para garantir que o nome seja informado corretamente.
	for {
		// Solicita o nome do produto a ser removido.
		fmt.Print("\nDigite o nome: ")
		scanner.Scan()
		nome = scanner.Text()

		// Valida se o nome foi informado.
		if nome == "" {
			// Caso o nome esteja vazio, solicita novamente.
			fmt.Println("\nFavor informar o nome corretamente.\n")
			continue
		}
		// Sai do loop quando o nome é válido.
		break
	}

	// Busca o produto pelo nome e remove.
	m.dao.Remover(m.dao.BuscarPorNome(nome).GetID())
	// Exibe mensagem de sucesso.
	fmt.Println("Produto removido com sucesso!")
}
