package main

import (
	"errors"
	"sync"
)

type Operacao string

const (
	Criar     Operacao = "CRIAR"
	Listar    Operacao = "LISTAR"
	Buscar    Operacao = "BUSCAR"
	Atualizar Operacao = "ATUALIZAR"
	Excluir   Operacao = "EXCLUIR"
)

type Requisicao struct {
	Operacao     Operacao
	ID           int
	Compromisso  Compromisso
	RespostaChan chan interface{}
}

type Gerenciador struct {
	mu           sync.Mutex
	compromissos map[int]Compromisso
	proximoID    int
	requisicoes  chan Requisicao
}

func NovoGerenciador() *Gerenciador {
	g := &Gerenciador{
		compromissos: make(map[int]Compromisso),
		proximoID:    1,
		requisicoes:  make(chan Requisicao),
	}

	go g.processarRequisicoes()

	return g
}

func (g *Gerenciador) processarRequisicoes() {
	for req := range g.requisicoes {
		switch req.Operacao {
		case Criar:
			g.mu.Lock()
			req.Compromisso.ID = g.proximoID
			g.compromissos[g.proximoID] = req.Compromisso
			g.proximoID++
			g.mu.Unlock()
			req.RespostaChan <- req.Compromisso
		case Listar:
			g.mu.Lock()
			lista := []Compromisso{}
			for _, c := range g.compromissos {
				lista = append(lista, c)
			}
			g.mu.Unlock()
			req.RespostaChan <- lista
		case Buscar:
			g.mu.Lock()
			comp, ok := g.compromissos[req.ID]
			g.mu.Unlock()
			if ok {
				req.RespostaChan <- comp
			} else {
				req.RespostaChan <- errors.New("compromisso não encontrado")
			}
		case Atualizar:
			g.mu.Lock()
			_, ok := g.compromissos[req.ID]
			if ok {
				req.Compromisso.ID = req.ID
				g.compromissos[req.ID] = req.Compromisso
				req.RespostaChan <- req.Compromisso
			} else {
				req.RespostaChan <- errors.New("compromisso não encontrado")
			}
			g.mu.Unlock()
		case Excluir:
			g.mu.Lock()
			_, ok := g.compromissos[req.ID]
			if ok {
				delete(g.compromissos, req.ID)
				req.RespostaChan <- nil
			} else {
				req.RespostaChan <- errors.New("compromisso não encontrado")
			}
			g.mu.Unlock()
		}
	}
}
