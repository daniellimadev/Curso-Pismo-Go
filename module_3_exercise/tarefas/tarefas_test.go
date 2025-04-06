package tarefas

import (
	"testing"
)

func TestAdicionarTarefa(t *testing.T) {
	tests := []struct {
		name        string
		descricao   string
		esperaErro  bool
		esperaID    TarefaID
		esperaTotal int
	}{
		{"Descrição válida", "Tarefa válida", false, 1, 1},
		{"Descrição vazia", "", true, 0, 0},
		{"Descrição com espaços", "   ", true, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NovaListaDeTarefas()
			tarefa, err := l.AdicionarTarefa(tt.descricao)

			if (err != nil) != tt.esperaErro {
				t.Errorf("AdicionarTarefa() erro = %v, esperaErro %v", err, tt.esperaErro)
				return
			}

			if !tt.esperaErro {
				if tarefa.ID != tt.esperaID {
					t.Errorf("AdicionarTarefa() ID = %v, espera %v", tarefa.ID, tt.esperaID)
				}

				if tarefa.Status != StatusPendente {
					t.Errorf("AdicionarTarefa() Status = %v, espera %v", tarefa.Status, StatusPendente)
				}

				if len(l.tarefas) != tt.esperaTotal {
					t.Errorf("Número de tarefas = %v, espera %v", len(l.tarefas), tt.esperaTotal)
				}
			}
		})
	}
}

func TestMarcarComoConcluida(t *testing.T) {
	tests := []struct {
		name         string
		id           TarefaID
		esperaErro   bool
		esperaStatus StatusTarefa
	}{
		{"Tarefa existente pendente", 1, false, StatusConcluida},
		{"Tarefa existente já concluída", 2, false, StatusConcluida},
		{"Tarefa não existente", 999, true, ""},
	}

	l := NovaListaDeTarefas()
	l.tarefas[1] = Tarefa{ID: 1, Descricao: "Tarefa 1", Status: StatusPendente}
	l.tarefas[2] = Tarefa{ID: 2, Descricao: "Tarefa 2", Status: StatusConcluida}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := l.MarcarComoConcluida(tt.id)

			if (err != nil) != tt.esperaErro {
				t.Errorf("MarcarComoConcluida() erro = %v, esperaErro %v", err, tt.esperaErro)
				return
			}

			if !tt.esperaErro {
				tarefa, _ := l.BuscarTarefaPorID(tt.id)
				if tarefa.Status != tt.esperaStatus {
					t.Errorf("MarcarComoConcluida() status = %v, espera %v", tarefa.Status, tt.esperaStatus)
				}
			}
		})
	}
}

func TestRemoverTarefa(t *testing.T) {
	tests := []struct {
		name       string
		id         TarefaID
		esperaErro bool
		esperaLen  int
	}{
		{"Tarefa existente", 1, false, 1},
		{"Tarefa não existente", 999, true, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar uma nova lista para cada teste
			l := NovaListaDeTarefas()
			l.tarefas[1] = Tarefa{ID: 1, Descricao: "Tarefa 1"}
			l.tarefas[2] = Tarefa{ID: 2, Descricao: "Tarefa 2"}

			err := l.RemoverTarefa(tt.id)

			if (err != nil) != tt.esperaErro {
				t.Errorf("RemoverTarefa() erro = %v, esperaErro %v", err, tt.esperaErro)
				return
			}

			if len(l.tarefas) != tt.esperaLen {
				t.Errorf("RemoverTarefa() len = %v, espera %v", len(l.tarefas), tt.esperaLen)
			}
		})
	}
}

func TestListarTarefas(t *testing.T) {
	tests := []struct {
		name       string
		filtro     StatusTarefa
		esperaLen  int
		esperaErro bool
	}{
		{"Todas as tarefas", "", 3, false},
		{"Tarefas pendentes", StatusPendente, 2, false},
		{"Tarefas concluídas", StatusConcluida, 1, false},
		{"Filtro inválido", "invalido", 0, true},
	}

	l := NovaListaDeTarefas()
	l.tarefas[1] = Tarefa{ID: 1, Descricao: "Tarefa 1", Status: StatusPendente}
	l.tarefas[2] = Tarefa{ID: 2, Descricao: "Tarefa 2", Status: StatusPendente}
	l.tarefas[3] = Tarefa{ID: 3, Descricao: "Tarefa 3", Status: StatusConcluida}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tarefas, err := l.ListarTarefas(tt.filtro)

			if (err != nil) != tt.esperaErro {
				t.Errorf("ListarTarefas() erro = %v, esperaErro %v", err, tt.esperaErro)
				return
			}

			if !tt.esperaErro && len(tarefas) != tt.esperaLen {
				t.Errorf("ListarTarefas() len = %v, espera %v", len(tarefas), tt.esperaLen)
			}

			// Verificar ordenação
			if !tt.esperaErro && len(tarefas) > 1 {
				for i := 1; i < len(tarefas); i++ {
					if tarefas[i-1].ID > tarefas[i].ID {
						t.Errorf("ListarTarefas() não está ordenado corretamente")
					}
				}
			}
		})
	}
}

func TestBuscarTarefaPorID(t *testing.T) {
	tests := []struct {
		name       string
		id         TarefaID
		esperaErro bool
	}{
		{"Tarefa existente", 1, false},
		{"Tarefa não existente", 999, true},
	}

	l := NovaListaDeTarefas()
	l.tarefas[1] = Tarefa{ID: 1, Descricao: "Tarefa 1"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := l.BuscarTarefaPorID(tt.id)

			if (err != nil) != tt.esperaErro {
				t.Errorf("BuscarTarefaPorID() erro = %v, esperaErro %v", err, tt.esperaErro)
			}
		})
	}
}
