package tarefas

import (
	"errors"
	"sort"
	"strings"
)

// Definindo tipos
type TarefaID int

type StatusTarefa string

const (
	StatusPendente  StatusTarefa = "pendente"
	StatusConcluida StatusTarefa = "concluida"
)

type Tarefa struct {
	ID        TarefaID
	Descricao string
	Status    StatusTarefa
}

// Definindo erros
var (
	ErroTarefaNaoEncontrada = errors.New("tarefa não encontrada")
	ErroDescricaoVazia      = errors.New("descrição da tarefa não pode ser vazia")
	ErroStatusInvalido      = errors.New("status de tarefa inválido")
)

// Interface GerenciadorTarefas
type GerenciadorTarefas interface {
	AdicionarTarefa(descricao string) (Tarefa, error)
	MarcarComoConcluida(id TarefaID) error
	RemoverTarefa(id TarefaID) error
	ListarTarefas(filtroStatus StatusTarefa) ([]Tarefa, error)
	BuscarTarefaPorID(id TarefaID) (Tarefa, error)
}

// Implementação ListaDeTarefas
type ListaDeTarefas struct {
	tarefas   map[TarefaID]Tarefa
	proximoID TarefaID
}

func NovaListaDeTarefas() *ListaDeTarefas {
	return &ListaDeTarefas{
		tarefas:   make(map[TarefaID]Tarefa),
		proximoID: 1,
	}
}

// Verificação de interface
var _ GerenciadorTarefas = (*ListaDeTarefas)(nil)

// Implementação dos métodos
func (l *ListaDeTarefas) AdicionarTarefa(descricao string) (Tarefa, error) {
	descricao = strings.TrimSpace(descricao)
	if descricao == "" {
		return Tarefa{}, ErroDescricaoVazia
	}

	tarefa := Tarefa{
		ID:        l.proximoID,
		Descricao: descricao,
		Status:    StatusPendente,
	}

	l.tarefas[l.proximoID] = tarefa
	l.proximoID++

	return tarefa, nil
}

func (l *ListaDeTarefas) MarcarComoConcluida(id TarefaID) error {
	tarefa, existe := l.tarefas[id]
	if !existe {
		return ErroTarefaNaoEncontrada
	}

	if tarefa.Status == StatusPendente {
		tarefa.Status = StatusConcluida
		l.tarefas[id] = tarefa
	}

	return nil
}

func (l *ListaDeTarefas) RemoverTarefa(id TarefaID) error {
	if _, existe := l.tarefas[id]; !existe {
		return ErroTarefaNaoEncontrada
	}

	delete(l.tarefas, id)
	return nil
}

func (l *ListaDeTarefas) ListarTarefas(filtroStatus StatusTarefa) ([]Tarefa, error) {
	var tarefas []Tarefa

	// Validar filtroStatus
	if filtroStatus != "" && filtroStatus != StatusPendente && filtroStatus != StatusConcluida {
		return nil, ErroStatusInvalido
	}

	for _, tarefa := range l.tarefas {
		if filtroStatus == "" || tarefa.Status == filtroStatus {
			tarefas = append(tarefas, tarefa)
		}
	}

	// Ordenar por ID
	sort.Slice(tarefas, func(i, j int) bool {
		return tarefas[i].ID < tarefas[j].ID
	})

	return tarefas, nil
}

func (l *ListaDeTarefas) BuscarTarefaPorID(id TarefaID) (Tarefa, error) {
	tarefa, existe := l.tarefas[id]
	if !existe {
		return Tarefa{}, ErroTarefaNaoEncontrada
	}

	return tarefa, nil
}
