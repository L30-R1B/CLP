package ui

import (
	"bufio"
	"fmt"
	"strconv"

	"clp-go-version/data"
	"clp-go-version/entidades"
)

// MenuVenda representa o menu para gerenciamento de vendas.
// Ele oferece opções para listar, adicionar e remover vendas do sistema.
type MenuVenda struct {
	daoVenda   *data.DAOVenda   // Acesso ao DAO para operações de vendas
	daoProduto *data.DAOProduto // Acesso ao DAO para operações de produtos
}

// NewMenuVenda cria uma nova instância de MenuVenda.
// Retorna um ponteiro para uma instância do menu de vendas com o DAO de vendas e produtos configurados.
func NewMenuVenda() *MenuVenda {
	return &MenuVenda{
		daoVenda:   data.GetVendaInstance(),
		daoProduto: data.GetInstance(),
	}
}

// MostrarTitulo exibe o título do menu de vendas.
// Este método imprime o título do menu para o usuário.
func (m *MenuVenda) MostrarTitulo() {
	fmt.Println("MENU VENDAS")
}

// MostrarOpcoes exibe as opções disponíveis no menu.
// As opções incluem listar, adicionar, remover ou voltar.
func (m *MenuVenda) MostrarOpcoes() {
	fmt.Println("0 -> VOLTAR")
	fmt.Println("1 -> LISTAR")
	fmt.Println("2 -> ADICIONAR")
	fmt.Println("3 -> REMOVER")
}

// MostrarMenu exibe o menu e gerencia as opções.
// Este método exibe o menu e processa as escolhas do usuário em um loop contínuo.
func (m *MenuVenda) MostrarMenu(scanner *bufio.Scanner) {
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
func (m *MenuVenda) ExecutarOpcao(opcao int, scanner *bufio.Scanner) int {
	switch opcao {
	case 0:
		// Caso 0: Retorna 0 para sair do menu.
		return 0
	case 1:
		// Caso 1: Chama o método Listar para exibir todas as vendas.
		m.Listar()
	case 2:
		// Caso 2: Chama o método Adicionar para adicionar uma nova venda.
		m.Adicionar(scanner)
	case 3:
		// Caso 3: Chama o método Remover para remover uma venda.
		m.Remover(scanner)
	default:
		// Caso inválido: Exibe mensagem de erro.
		fmt.Println("OPÇÃO INVÁLIDA\n")
	}
	// Retorna 1 para continuar exibindo o menu.
	return 1
}

// Listar exibe todas as vendas cadastradas no sistema.
// Este método imprime a lista de vendas do DAO.
func (m *MenuVenda) Listar() {
	// Exibe os dados do DAO de vendas.
	fmt.Println(m.daoVenda.String())
}

// Adicionar adiciona uma nova venda ao sistema.
// O método solicita ao usuário o nome do produto e a quantidade, e valida as entradas.
func (m *MenuVenda) Adicionar(scanner *bufio.Scanner) {
	venda := entidades.NewVenda() // Cria uma nova venda

	// Loop para adicionar itens à venda.
	for {
		var produto *entidades.Produto
		var qtd int

		// Solicita ao usuário o nome do produto a ser adicionado.
		for {
			fmt.Print("\nDigite o nome do produto: ")
			scanner.Scan()
			nomeProduto := scanner.Text()

			// Busca o produto pelo nome.
			produto = m.daoProduto.BuscarPorNome(nomeProduto)
			if produto == nil {
				// Caso o produto não seja encontrado, solicita novamente.
				fmt.Println("Produto não encontrado. Tente novamente.")
				continue
			}

			// Solicita a quantidade do produto.
			fmt.Print("Digite a quantidade: ")
			scanner.Scan()
			qtd, _ = strconv.Atoi(scanner.Text())

			// Valida se a quantidade é válida.
			if qtd <= 0 {
				fmt.Println("Quantidade inválida. Tente novamente.")
				continue
			}
			break
		}

		// Adiciona o produto à venda.
		venda.AdicionarItem(*produto, qtd)

		// Pergunta se o usuário deseja adicionar outro produto à venda.
		fmt.Print("\nDeseja adicionar outro produto à venda (1-SIM/0-NAO)? ")
		scanner.Scan()
		opcao, _ := strconv.Atoi(scanner.Text())
		if opcao != 1 {
			// Sai do loop se o usuário não quiser adicionar mais produtos.
			break
		}
	}

	// Exibe a nota fiscal com os itens da venda.
	fmt.Println("\n\nNOTA FISCAL\n", venda.String())
	// Adiciona a venda ao DAO.
	m.daoVenda.Adicionar(venda)
}

// Remover remove uma venda com base no ID.
// O método solicita ao usuário o ID da venda a ser removida e, caso encontrado, a remove do sistema.
func (m *MenuVenda) Remover(scanner *bufio.Scanner) {
	var id int64

	// Loop para garantir que o ID da venda seja informado corretamente.
	for {
		// Solicita o ID da venda.
		fmt.Print("\nDigite o id: ")
		scanner.Scan()
		id, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		// Valida se o ID é válido.
		if id <= 0 {
			fmt.Println("ID inválido. Tente novamente.")
			continue
		}
		// Sai do loop quando o ID for válido.
		break
	}

	// Remove a venda com o ID informado.
	m.daoVenda.Remover(id)
	// Exibe mensagem de sucesso.
	fmt.Println("Venda removida com sucesso.")
}
