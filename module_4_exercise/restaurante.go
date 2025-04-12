package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	gerenciador := NovoGerenciadorPedidos()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		exibirMenu()
		opcao, err := lerOpcao(scanner)
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		if opcao == 4 {
			fmt.Println("Saindo...")
			break
		}

		err = executarOpcao(opcao, gerenciador, scanner)
		if err != nil {
			fmt.Println("Erro:", err)
		}
	}
}

func exibirMenu() {
	fmt.Println("\nMenu Principal:")
	fmt.Println("1. Adicionar um novo pedido")
	fmt.Println("2. Listar todos os pedidos pendentes")
	fmt.Println("3. Marcar um pedido como entregue")
	fmt.Println("4. Sair")
	fmt.Print("Escolha uma opção: ")
}

func lerOpcao(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("falha ao ler entrada")
	}

	texto := strings.TrimSpace(scanner.Text())
	opcao, err := strconv.Atoi(texto)
	if err != nil {
		return 0, fmt.Errorf("opção inválida")
	}

	if opcao < 1 || opcao > 4 {
		return 0, fmt.Errorf("opção deve ser entre 1 e 4")
	}

	return opcao, nil
}

func executarOpcao(opcao int, gerenciador *GerenciadorPedidos, scanner *bufio.Scanner) error {
	switch opcao {
	case 1:
		return adicionarNovoPedido(gerenciador, scanner)
	case 2:
		return listarPedidos(gerenciador)
	case 3:
		return marcarPedidoComoEntregue(gerenciador, scanner)
	default:
		return nil
	}
}

func adicionarNovoPedido(gerenciador *GerenciadorPedidos, scanner *bufio.Scanner) error {
	fmt.Print("Digite a descrição do pedido: ")
	if !scanner.Scan() {
		return fmt.Errorf("falha ao ler entrada")
	}

	descricao := strings.TrimSpace(scanner.Text())
	pedido, err := gerenciador.AdicionarPedido(descricao)
	if err != nil {
		return err
	}

	fmt.Printf("Pedido #%d adicionado com sucesso!\n", pedido.Numero)
	return nil
}

func listarPedidos(gerenciador *GerenciadorPedidos) error {
	pedidos := gerenciador.ListarPedidos()
	if len(pedidos) == 0 {
		fmt.Println("Nenhum pedido cadastrado.")
		return nil
	}

	fmt.Println("\nLista de Pedidos:")
	for _, pedido := range pedidos {
		fmt.Println(pedido)
	}
	return nil
}

func marcarPedidoComoEntregue(gerenciador *GerenciadorPedidos, scanner *bufio.Scanner) error {
	fmt.Print("Digite o número do pedido a marcar como entregue: ")
	if !scanner.Scan() {
		return fmt.Errorf("falha ao ler entrada")
	}

	texto := strings.TrimSpace(scanner.Text())
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return fmt.Errorf("número do pedido inválido")
	}

	err = gerenciador.MarcarComoEntregue(numero)
	if err != nil {
		return err
	}

	fmt.Printf("Pedido #%d marcado como entregue!\n", numero)
	return nil
}
