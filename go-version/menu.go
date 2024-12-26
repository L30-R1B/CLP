package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Menu interface {
	ShowMenu()
}

type MenuAbstrato struct{}

func (m MenuAbstrato) MostrarTitulo(titulo string) {
	fmt.Println("\n", titulo)
}

func (m MenuAbstrato) MostrarOpcoes(opcoes []string) {
	for i, opcao := range opcoes {
		fmt.Printf("%d -> %s\n", i, opcao)
	}
}

type MainMenu struct {
	productMenu *ProductMenu
	saleMenu    *SaleMenu
}

func NewMainMenu() *MainMenu {
	return &MainMenu{
		productMenu: NewProductMenu(),
		saleMenu:    NewSaleMenu(),
	}
}

func (m *MainMenu) ShowMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMENU PRINCIPAL")
		fmt.Println("0 -> FECHAR PROGRAMA")
		fmt.Println("1 -> PRODUTO")
		fmt.Println("2 -> VENDA")

		fmt.Print("INFORME A SUA OPÇÃO: ")
		if !scanner.Scan() {
			break
		}
		option := scanner.Text()

		switch option {
		case "0":
			fmt.Println("Encerrando o programa.")
			return
		case "1":
			m.productMenu.ShowMenu()
		case "2":
			m.saleMenu.ShowMenu()
		default:
			fmt.Println("OPÇÃO INVÁLIDA\n")
		}
	}
}

type ProductMenu struct {
	dao *DAOProduto
}

func NewProductMenu() *ProductMenu {
	return &ProductMenu{
		dao: GetProdutoInstance(),
	}
}

func (p *ProductMenu) ShowMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMENU PRODUTOS")
		fmt.Println("0 -> VOLTAR")
		fmt.Println("1 -> LISTAR")
		fmt.Println("2 -> ADICIONAR")
		fmt.Println("3 -> REMOVER")

		fmt.Print("INFORME A SUA OPÇÃO: ")
		if !scanner.Scan() {
			break
		}
		option := scanner.Text()

		switch option {
		case "0":
			return
		case "1":
			p.listProducts()
		case "2":
			p.addProduct(scanner)
		case "3":
			p.removeProduct(scanner)
		default:
			fmt.Println("OPÇÃO INVÁLIDA\n")
		}
	}
}

func (p *ProductMenu) listProducts() {
	fmt.Println("\nLISTA DE PRODUTOS:")
	for _, produto := range p.dao.Listar() {
		fmt.Println(produto.String())
	}
}

func (p *ProductMenu) addProduct(scanner *bufio.Scanner) {
	var name string
	var price float64
	var err error

	for {
		fmt.Print("\nDigite o nome: ")
		if !scanner.Scan() {
			break
		}
		name = scanner.Text()

		fmt.Print("Digite o valor: ")
		if !scanner.Scan() {
			break
		}
		price, err = strconv.ParseFloat(scanner.Text(), 64)

		if name == "" || price <= 0.0 || err != nil {
			fmt.Println("\nFavor informar os dados corretamente.")
		} else {
			break
		}
	}

	p.dao.Adicionar(NewProduto(name, price))
	fmt.Println("Produto adicionado com sucesso!")
}

func (p *ProductMenu) removeProduct(scanner *bufio.Scanner) {
	var name string

	for {
		fmt.Print("\nDigite o nome do produto: ")
		if !scanner.Scan() {
			break
		}
		name = scanner.Text()

		if name == "" {
			fmt.Println("\nFavor informar o nome corretamente.")
		} else {
			break
		}
	}

	p.dao.RemoverPorNome(name)
	fmt.Println("Produto removido com sucesso!")
}

type SaleMenu struct {
	daoVenda   *DAOVenda
	daoProduto *DAOProduto
}

func NewSaleMenu() *SaleMenu {
	return &SaleMenu{
		daoVenda:   GetVendaInstance(),
		daoProduto: GetProdutoInstance(),
	}
}

func (s *SaleMenu) ShowMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMENU VENDAS")
		fmt.Println("0 -> VOLTAR")
		fmt.Println("1 -> LISTAR")
		fmt.Println("2 -> ADICIONAR")
		fmt.Println("3 -> REMOVER")

		fmt.Print("INFORME A SUA OPÇÃO: ")
		if !scanner.Scan() {
			break
		}
		option := scanner.Text()

		switch option {
		case "0":
			return
		case "1":
			s.listSales()
		case "2":
			s.addSale(scanner)
		case "3":
			s.removeSale(scanner)
		default:
			fmt.Println("OPÇÃO INVÁLIDA\n")
		}
	}
}

func (s *SaleMenu) listSales() {
	fmt.Println("\nLISTA DE VENDAS:")
	for _, venda := range s.daoVenda.Listar() {
		fmt.Println(venda.String())
	}
}

func (s *SaleMenu) addSale(scanner *bufio.Scanner) {
	venda := NewVenda()

	for {
		fmt.Print("\nDigite o nome do produto: ")
		if !scanner.Scan() {
			break
		}
		productName := scanner.Text()

		produto := s.daoProduto.BuscarPorNome(productName)
		if produto == nil {
			fmt.Println("Produto não encontrado!")
			continue
		}

		fmt.Print("Digite a quantidade: ")
		if !scanner.Scan() {
			break
		}
		qtd, err := strconv.Atoi(scanner.Text())

		if err != nil || qtd <= 0 {
			fmt.Println("Quantidade inválida!")
			continue
		}

		venda.AdicionarItem(*produto, qtd)
		break
	}

	s.daoVenda.Adicionar(venda)
	fmt.Println("Venda adicionada com sucesso!")
}

func (s *SaleMenu) removeSale(scanner *bufio.Scanner) {
	var id int64

	for {
		fmt.Print("\nDigite o ID da venda para remover: ")
		if !scanner.Scan() {
			break
		}
		id, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		if id <= 0 {
			fmt.Println("\nID inválido.")
		} else {
			break
		}
	}

	s.daoVenda.RemoverPorID(id)
	fmt.Println("Venda removida com sucesso!")
}
