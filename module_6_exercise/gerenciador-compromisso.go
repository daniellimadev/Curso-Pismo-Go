package main

import (
	"errors"
	"sync"
)

// Operacao define os tipos de operações que podem ser executadas pelo gerenciador
type Operacao string

const (
	Criar     Operacao = "CRIAR"     // Operação para criar um novo compromisso
	Listar    Operacao = "LISTAR"    // Operação para listar todos os compromissos
	Buscar    Operacao = "BUSCAR"    // Operação para buscar um compromisso pelo ID
	Atualizar Operacao = "ATUALIZAR" // Operação para atualizar um compromisso existente
	Excluir   Operacao = "EXCLUIR"   // Operação para excluir um compromisso
)

// Requisicao representa uma solicitação ao gerenciador de compromissos
type Requisicao struct {
	Operacao     Operacao         // Tipo da operação a ser executada
	ID           int              // ID do compromisso (para buscar, atualizar ou excluir)
	Compromisso  Compromisso      // Dados do compromisso (para criar ou atualizar)
	RespostaChan chan interface{} // Canal por onde será enviada a resposta
}

// Gerenciador é responsável por armazenar e manipular compromissos de forma concorrente
type Gerenciador struct {
	mu           sync.Mutex          // Mutex para garantir segurança em operações concorrentes
	compromissos map[int]Compromisso // Mapa que armazena os compromissos usando o ID como chave
	proximoID    int                 // Próximo ID disponível para novo compromisso
	requisicoes  chan Requisicao     // Canal de entrada para requisições
}

// NovoGerenciador cria e inicializa um novo gerenciador e inicia o processamento de requisições
func NovoGerenciador() *Gerenciador {
	g := &Gerenciador{
		compromissos: make(map[int]Compromisso),
		proximoID:    1,
		requisicoes:  make(chan Requisicao),
	}

	// Inicia a goroutine responsável por processar as requisições recebidas
	go g.processarRequisicoes()

	return g
}

// processarRequisicoes escuta o canal de requisições e executa a operação correspondente
func (g *Gerenciador) processarRequisicoes() {
	for req := range g.requisicoes {
		switch req.Operacao {

		case Criar:
			// Criação de compromisso com ID único
			g.mu.Lock()
			req.Compromisso.ID = g.proximoID
			g.compromissos[g.proximoID] = req.Compromisso
			g.proximoID++
			g.mu.Unlock()
			req.RespostaChan <- req.Compromisso

		case Listar:
			// Lista todos os compromissos
			g.mu.Lock()
			lista := []Compromisso{}
			for _, c := range g.compromissos {
				lista = append(lista, c)
			}
			g.mu.Unlock()
			req.RespostaChan <- lista

		case Buscar:
			// Busca um compromisso pelo ID
			g.mu.Lock()
			comp, ok := g.compromissos[req.ID]
			g.mu.Unlock()
			if ok {
				req.RespostaChan <- comp
			} else {
				req.RespostaChan <- errors.New("compromisso não encontrado")
			}

		case Atualizar:
			// Atualiza um compromisso existente
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
			// Exclui um compromisso pelo ID
			g.mu.Lock()
			_, ok := g.compromissos[req.ID]
			if ok {
				delete(g.compromissos, req.ID)
				req.RespostaChan <- nil // nil indica sucesso sem resposta
			} else {
				req.RespostaChan <- errors.New("compromisso não encontrado")
			}
			g.mu.Unlock()
		}
	}
}
