package main

import (
	"fmt"
	"sync"
	"time"
)

type GerenciadorPedidos struct {
	pedidos   []*Pedido
	proximoID int
	mu        sync.Mutex
}

func NovoGerenciadorPedidos() *GerenciadorPedidos {
	return &GerenciadorPedidos{
		pedidos:   make([]*Pedido, 0),
		proximoID: 1,
	}
}

func (g *GerenciadorPedidos) AdicionarPedido(descricao string) (*Pedido, error) {
	if descricao == "" {
		return nil, fmt.Errorf("descrição do pedido é obrigatória")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	pedido := NovoPedido(g.proximoID, descricao)
	g.pedidos = append(g.pedidos, pedido)
	g.proximoID++

	// Inicia a simulação de preparo em uma goroutine separada
	go g.simularPreparo(pedido)

	return pedido, nil
}

func (g *GerenciadorPedidos) simularPreparo(pedido *Pedido) {
	time.Sleep(5 * time.Second)

	g.mu.Lock()
	defer g.mu.Unlock()

	for _, p := range g.pedidos {
		if p.Numero == pedido.Numero {
			p.MarcarComoEntregue()
			fmt.Printf("Pedido #%d entregue automaticamente após preparo!\n", p.Numero)
			break
		}
	}
}

func (g *GerenciadorPedidos) ListarPedidos() []*Pedido {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Retorna uma cópia para evitar problemas de concorrência
	result := make([]*Pedido, len(g.pedidos))
	copy(result, g.pedidos)
	return result
}

func (g *GerenciadorPedidos) MarcarComoEntregue(numero int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, pedido := range g.pedidos {
		if pedido.Numero == numero {
			pedido.MarcarComoEntregue()
			return nil
		}
	}

	return fmt.Errorf("pedido %d não encontrado", numero)
}

func (g *GerenciadorPedidos) ObterPedido(numero int) (*Pedido, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, pedido := range g.pedidos {
		if pedido.Numero == numero {
			return pedido, nil
		}
	}

	return nil, fmt.Errorf("pedido %d não encontrado", numero)
}
