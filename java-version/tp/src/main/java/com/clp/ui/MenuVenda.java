package com.clp.ui;

import com.clp.data.DAOProduto;
import com.clp.data.DAOVenda;
import com.clp.entidades.Produto;
import com.clp.entidades.Venda;
import java.util.Scanner;

public class MenuVenda extends MenuEntidade {

  private DAOVenda daoVenda;
  private DAOProduto daoProduto;

  public MenuVenda() {
    super();
    daoVenda = DAOVenda.getInstance();
    daoProduto = DAOProduto.getInstance();
  }

  @Override
  protected void mostrarTitulo() {
    System.out.println("MENU VENDAS");
  }

  @Override
  protected void listar() {
    System.out.println(daoVenda.toString());
  }

  @Override
  protected void adicionar(Scanner scanner) {

    Venda venda = new Venda();
    Produto produto = null;
    int qtd = 0;

    while (true) {

      while (true) {

        try {

          System.out.print("\nDigite o nome do produto: ");
          produto = daoProduto.buscar(scanner.next());

          System.out.print("Digite a quantidade: ");
          qtd = scanner.nextInt();

          if (produto == null || qtd <= 0)
            throw new Exception("\nFavor informar os dados corretamente.\n");
          else
            break;

        } catch (Exception ex) {
          System.err.println(ex.getMessage());
        }
      }

      venda.adicionarItem(produto, qtd);

      System.out.print("\nDeseja adicionar outro produto à venda (1-SIM/0-NAO)? ");

      if (scanner.nextInt() != 1)
        break;
    }

    System.out.println("\n\nNOTA FISCAL\n" + venda.toString());

    daoVenda.adicionar(venda);
  }

  @Override
  protected void remover(Scanner scanner) {

    long id = 0;

    while (true) {

      try {

        System.out.print("\nDigite o id: ");
        id = scanner.nextLong();

        if (id <= 0.0)
          throw new Exception("\nFavor informar os dados corretamente.\n");
        else
          break;

      } catch (Exception ex) {
        System.err.println(ex.getMessage());
      }
    }
    daoVenda.remover(id);
  }

}
