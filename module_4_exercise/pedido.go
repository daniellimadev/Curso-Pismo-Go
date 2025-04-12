package main

import (
	"fmt"
	"time"
)

type Status string

const (
	Pendente Status = "pendente"
	Entregue Status = "entregue"
)

type Pedido struct {
	Numero    int
	Descricao string
	Status    Status
	CriadoEm  time.Time
}

func NovoPedido(numero int, descricao string) *Pedido {
	return &Pedido{
		Numero:    numero,
		Descricao: descricao,
		Status:    Pendente,
		CriadoEm:  time.Now(),
	}
}

func (p *Pedido) MarcarComoEntregue() {
	p.Status = Entregue
}

func (p *Pedido) String() string {
	return fmt.Sprintf("Pedido #%d - %s (%s)", p.Numero, p.Descricao, p.Status)
}
