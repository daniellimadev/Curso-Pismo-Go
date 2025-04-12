package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// Redireciona a saída padrão para capturar o que é impresso
func capturarSaida(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestExibirMenu(t *testing.T) {
	saida := capturarSaida(exibirMenu)

	// Note o \n no início para corresponder à saída real
	esperado := "\nMenu Principal:\n1. Adicionar um novo pedido\n2. Listar todos os pedidos pendentes\n3. Marcar um pedido como entregue\n4. Sair\nEscolha uma opção: "

	if saida != esperado {
		t.Errorf("exibirMenu() = %q, esperado %q", saida, esperado)
	}
}

func TestLerOpcao(t *testing.T) {
	tests := []struct {
		nome     string
		input    string
		esperado int
		erro     string
	}{
		{"Opção válida", "1\n", 1, ""},
		{"Opção com espaços", "   2   \n", 2, ""},
		{"Opção inválida (letra)", "a\n", 0, "opção inválida"},
		{"Opção inválida (fora do range)", "5\n", 0, "opção deve ser entre 1 e 4"},
		{"Input vazio", "\n", 0, "opção inválida"},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			opcao, err := lerOpcao(scanner)

			if err != nil && err.Error() != tt.erro {
				t.Errorf("lerOpcao() erro = %v, esperado %v", err, tt.erro)
			}
			if opcao != tt.esperado {
				t.Errorf("lerOpcao() = %v, esperado %v", opcao, tt.esperado)
			}
		})
	}
}

func TestAdicionarNovoPedido(t *testing.T) {
	tests := []struct {
		nome      string
		input     string
		esperado  string
		erro      string
		descricao string
	}{
		{
			nome:      "Descrição válida",
			input:     "Hambúrguer\n",
			esperado:  "Pedido #1 adicionado com sucesso!\n",
			erro:      "",
			descricao: "Hambúrguer",
		},
		{
			nome:      "Descrição vazia",
			input:     "\n",
			esperado:  "",
			erro:      "descrição do pedido é obrigatória", // Corrigido aqui
			descricao: "",
		},
		{
			nome:      "Descrição com espaços",
			input:     "   Pizza Margherita   \n",
			esperado:  "Pedido #1 adicionado com sucesso!\n",
			erro:      "",
			descricao: "Pizza Margherita",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			g := NovoGerenciadorPedidos()
			scanner := bufio.NewScanner(strings.NewReader(tt.input))

			saida := capturarSaida(func() {
				err := adicionarNovoPedido(g, scanner)
				if err != nil {
					if tt.erro == "" {
						t.Errorf("Erro inesperado: %v", err)
					} else if !strings.Contains(err.Error(), tt.erro) {
						t.Errorf("Erro = %v, esperado conter %q", err, tt.erro)
					}
					return
				}
				if tt.erro != "" {
					t.Errorf("Esperado erro contendo %q, mas nenhum erro ocorreu", tt.erro)
				}
			})

			if !strings.Contains(saida, tt.esperado) {
				t.Errorf("Saída = %q, esperado conter %q", saida, tt.esperado)
			}

			if tt.descricao != "" {
				pedidos := g.ListarPedidos()
				if len(pedidos) == 0 {
					t.Error("Nenhum pedido foi adicionado")
				} else if pedidos[0].Descricao != strings.TrimSpace(tt.descricao) {
					t.Errorf("Descrição do pedido = %q, esperado %q",
						pedidos[0].Descricao, strings.TrimSpace(tt.descricao))
				}
			}
		})
	}
}

func TestListarPedidoss(t *testing.T) {
	tests := []struct {
		nome     string
		pedidos  []*Pedido
		esperado string
	}{
		{"Lista vazia", []*Pedido{}, "Nenhum pedido cadastrado.\n"},
		{"Um pedido", []*Pedido{NovoPedido(1, "Salada")}, "Pedido #1 - Salada (pendente)\n"},
		{"Múltiplos pedidos", []*Pedido{
			NovoPedido(1, "Sopa"),
			NovoPedido(2, "Lasanha"),
		}, "Pedido #1 - Sopa (pendente)\nPedido #2 - Lasanha (pendente)\n"},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			g := NovoGerenciadorPedidos()
			// Adiciona pedidos ao gerenciador (usando reflexão ou exportando campos se necessário)
			for _, p := range tt.pedidos {
				g.AdicionarPedido(p.Descricao)
			}

			saida := capturarSaida(func() {
				err := listarPedidos(g)
				if err != nil {
					t.Errorf("Erro inesperado: %v", err)
				}
			})

			if !strings.Contains(saida, tt.esperado) {
				t.Errorf("Saída = %q, esperado conter %q", saida, tt.esperado)
			}
		})
	}
}

func TestMarcarPedidoComoEntregue(t *testing.T) {
	tests := []struct {
		nome     string
		input    string
		setup    func(*GerenciadorPedidos)
		esperado string
		erro     string
	}{
		{
			"Pedido existente",
			"1\n",
			func(g *GerenciadorPedidos) { g.AdicionarPedido("Pizza") },
			"Pedido #1 marcado como entregue!\n",
			"",
		},
		{
			"Pedido inexistente",
			"999\n",
			func(g *GerenciadorPedidos) { g.AdicionarPedido("Sushi") },
			"",
			"pedido 999 não encontrado",
		},
		{
			"Entrada inválida",
			"abc\n",
			func(g *GerenciadorPedidos) {},
			"",
			"número do pedido inválido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			g := NovoGerenciadorPedidos()
			tt.setup(g)

			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			saida := capturarSaida(func() {
				err := marcarPedidoComoEntregue(g, scanner)
				if err != nil && tt.erro == "" {
					t.Errorf("Erro inesperado: %v", err)
				}
				if err != nil && !strings.Contains(err.Error(), tt.erro) {
					t.Errorf("Erro = %v, esperado conter %q", err, tt.erro)
				}
			})

			if !strings.Contains(saida, tt.esperado) {
				t.Errorf("Saída = %q, esperado conter %q", saida, tt.esperado)
			}
		})
	}
}

func TestExecutarOpcao(t *testing.T) {
	tests := []struct {
		nome     string
		opcao    int
		input    string
		esperado string
		erro     string
	}{
		{
			"Opção 1 - Adicionar pedido",
			1,
			"Hambúrguer\n",
			"Pedido #1 adicionado com sucesso!\n",
			"",
		},
		{
			"Opção 2 - Listar pedidos",
			2,
			"",
			"Nenhum pedido cadastrado.\n",
			"",
		},
		{
			"Opção 3 - Marcar como entregue (inválido)",
			3,
			"999\n",
			"",
			"pedido 999 não encontrado",
		},
		{
			"Opção inválida",
			99,
			"",
			"",
			"opção inválida",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			g := NovoGerenciadorPedidos()
			scanner := bufio.NewScanner(strings.NewReader(tt.input))

			saida := capturarSaida(func() {
				err := executarOpcao(tt.opcao, g, scanner)
				if err != nil && tt.erro == "" {
					t.Errorf("Erro inesperado: %v", err)
				}
				if err != nil && !strings.Contains(err.Error(), tt.erro) {
					t.Errorf("Erro = %v, esperado conter %q", err, tt.erro)
				}
			})

			if !strings.Contains(saida, tt.esperado) {
				t.Errorf("Saída = %q, esperado conter %q", saida, tt.esperado)
			}
		})
	}
}

func TestNovoPedido(t *testing.T) {
	pedido := NovoPedido(1, "Hambúrguer")
	if pedido.Numero != 1 {
		t.Errorf("Número do pedido esperado: 1, obtido: %d", pedido.Numero)
	}
	if pedido.Descricao != "Hambúrguer" {
		t.Errorf("Descrição do pedido esperada: 'Hambúrguer', obtida: '%s'", pedido.Descricao)
	}
	if pedido.Status != Pendente {
		t.Errorf("Status do pedido esperado: 'pendente', obtido: '%s'", pedido.Status)
	}
}

func TestAdicionarPedido(t *testing.T) {
	g := NovoGerenciadorPedidos()

	// Teste de adição válida
	pedido, err := g.AdicionarPedido("Pizza")
	if err != nil {
		t.Errorf("Erro inesperado ao adicionar pedido: %v", err)
	}
	if pedido.Numero != 1 {
		t.Errorf("Número do pedido esperado: 1, obtido: %d", pedido.Numero)
	}

	// Teste de descrição vazia
	_, err = g.AdicionarPedido("")
	if err == nil {
		t.Error("Esperado erro ao adicionar pedido com descrição vazia")
	}
}

func TestListarPedidos(t *testing.T) {
	g := NovoGerenciadorPedidos()

	// Lista vazia
	pedidos := g.ListarPedidos()
	if len(pedidos) != 0 {
		t.Errorf("Esperado 0 pedidos, obtido %d", len(pedidos))
	}

	// Adiciona pedidos e verifica
	g.AdicionarPedido("Salada")
	g.AdicionarPedido("Sopa")
	pedidos = g.ListarPedidos()
	if len(pedidos) != 2 {
		t.Errorf("Esperado 2 pedidos, obtido %d", len(pedidos))
	}
}

func TestMarcarComoEntregue(t *testing.T) {
	g := NovoGerenciadorPedidos()
	g.AdicionarPedido("Lasanha")

	// Teste de marcação válida
	err := g.MarcarComoEntregue(1)
	if err != nil {
		t.Errorf("Erro inesperado ao marcar pedido como entregue: %v", err)
	}

	pedido, _ := g.ObterPedido(1)
	if pedido.Status != Entregue {
		t.Errorf("Status do pedido esperado: 'entregue', obtido: '%s'", pedido.Status)
	}

	// Teste de pedido inexistente
	err = g.MarcarComoEntregue(999)
	if err == nil {
		t.Error("Esperado erro ao marcar pedido inexistente como entregue")
	}
}

func TestSimulacaoPreparo(t *testing.T) {
	g := NovoGerenciadorPedidos()
	pedido, _ := g.AdicionarPedido("Macarrão")

	if pedido.Status != Pendente {
		t.Errorf("Status inicial do pedido deve ser 'pendente'")
	}

	// Aguarda o tempo de preparo + um pouco mais
	time.Sleep(6 * time.Second)

	pedidoAtualizado, _ := g.ObterPedido(pedido.Numero)
	if pedidoAtualizado.Status != Entregue {
		t.Errorf("Após preparo, status do pedido deve ser 'entregue'")
	}
}

func TestConcorrencia(t *testing.T) {
	g := NovoGerenciadorPedidos()

	// Adiciona vários pedidos concorrentemente
	for i := 0; i < 100; i++ {
		go func(i int) {
			g.AdicionarPedido(fmt.Sprintf("Pedido %d", i))
		}(i)
	}

	// Marca vários pedidos como entregues concorrentemente
	for i := 0; i < 100; i++ {
		go func(i int) {
			g.MarcarComoEntregue(i + 1)
		}(i)
	}

	// Aguarda um tempo para todas as goroutines terminarem
	time.Sleep(2 * time.Second)

	// Verifica se todos os pedidos foram processados corretamente
	pedidos := g.ListarPedidos()
	if len(pedidos) != 100 {
		t.Errorf("Esperado 100 pedidos, obtido %d", len(pedidos))
	}
}
