package com.clp;

import com.clp.ui.MenuPrincipal;
import java.util.Scanner;

public class Main {

  public static void main(String[] args) {

    Scanner scanner = new Scanner(System.in);
    scanner.useDelimiter("\\R");

    new MenuPrincipal().mostrarMenu(scanner);

    scanner.close();
  }

}