package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tarefas/tarefas"
)

const (
	MenuPrompt = `
=== Gerenciador de Tarefas ===
1. Adicionar Tarefa
2. Marcar Tarefa como Concluída
3. Remover Tarefa
4. Listar Tarefas Pendentes
5. Listar Todas as Tarefas
6. Buscar Tarefa por ID
7. Sair
Escolha uma opção: `

	MsgSaindo             = "Saindo do programa..."
	MsgOpcaoInvalida      = "Opção inválida. Tente novamente."
	MsgErroEntradaID      = "ID inválido."
	MsgTarefaAdicionada   = "Tarefa (ID: %d) adicionada com sucesso."
	MsgTarefaConcluida    = "Tarefa (ID: %d) marcada como concluída."
	MsgTarefaRemovida     = "Tarefa (ID: %d) removida com sucesso."
	MsgListandoTarefas    = "--- Tarefas %s ---"
	MsgNenhumaTarefa      = "Nenhuma tarefa encontrada."
	MsgTarefaDetalhe      = "ID: %d | Status: %s | Descrição: %s"
	MsgErroEntradaDesc    = "Descrição inválida."
	MsgErroMarcarConcluir = "Erro ao marcar tarefa como concluída: %v"
	MsgErroRemover        = "Erro ao remover tarefa: %v"
	MsgErroBuscar         = "Erro ao buscar tarefa: %v"
)

func main() {
	gerenciador := tarefas.NovaListaDeTarefas()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(MenuPrompt)
		scanner.Scan()
		opcao := strings.TrimSpace(scanner.Text())

		switch opcao {
		case "1":
			handleAdicionarTarefa(gerenciador, scanner)
		case "2":
			handleMarcarConcluida(gerenciador, scanner)
		case "3":
			handleRemoverTarefa(gerenciador, scanner)
		case "4":
			handleListarTarefas(gerenciador, tarefas.StatusPendente)
		case "5":
			handleListarTarefas(gerenciador, "")
		case "6":
			handleBuscarTarefa(gerenciador, scanner)
		case "7":
			fmt.Println(MsgSaindo)
			return
		default:
			fmt.Println(MsgOpcaoInvalida)
		}
	}
}

func handleAdicionarTarefa(g tarefas.GerenciadorTarefas, scanner *bufio.Scanner) {
	fmt.Print("Digite a descrição da tarefa: ")
	scanner.Scan()
	descricao := scanner.Text()

	tarefa, err := g.AdicionarTarefa(descricao)
	if err != nil {
		fmt.Printf("%s Detalhe: %v\n", MsgErroEntradaDesc, err)
		return
	}

	fmt.Printf(MsgTarefaAdicionada+"\n", tarefa.ID)
}

func handleMarcarConcluida(g tarefas.GerenciadorTarefas, scanner *bufio.Scanner) {
	id, err := lerID(scanner)
	if err != nil {
		fmt.Printf("%s Detalhe: %v\n", MsgErroEntradaID, err)
		return
	}

	err = g.MarcarComoConcluida(id)
	if err != nil {
		fmt.Printf(MsgErroMarcarConcluir+"\n", err)
		return
	}

	fmt.Printf(MsgTarefaConcluida+"\n", id)
}

func handleRemoverTarefa(g tarefas.GerenciadorTarefas, scanner *bufio.Scanner) {
	id, err := lerID(scanner)
	if err != nil {
		fmt.Printf("%s Detalhe: %v\n", MsgErroEntradaID, err)
		return
	}

	err = g.RemoverTarefa(id)
	if err != nil {
		fmt.Printf(MsgErroRemover+"\n", err)
		return
	}

	fmt.Printf(MsgTarefaRemovida+"\n", id)
}

func handleListarTarefas(g tarefas.GerenciadorTarefas, filtro tarefas.StatusTarefa) {
	var titulo string
	if filtro == "" {
		titulo = "Todas"
	} else {
		titulo = string(filtro)
	}

	fmt.Printf(MsgListandoTarefas+"\n", titulo)

	tarefas, err := g.ListarTarefas(filtro)
	if err != nil {
		fmt.Printf("Erro ao listar tarefas: %v\n", err)
		return
	}

	if len(tarefas) == 0 {
		fmt.Println(MsgNenhumaTarefa)
		return
	}

	for _, t := range tarefas {
		fmt.Printf(MsgTarefaDetalhe+"\n", t.ID, t.Status, t.Descricao)
	}
}

func handleBuscarTarefa(g tarefas.GerenciadorTarefas, scanner *bufio.Scanner) {
	id, err := lerID(scanner)
	if err != nil {
		fmt.Printf("%s Detalhe: %v\n", MsgErroEntradaID, err)
		return
	}

	tarefa, err := g.BuscarTarefaPorID(id)
	if err != nil {
		fmt.Printf(MsgErroBuscar+"\n", err)
		return
	}

	fmt.Println("--- Tarefa Encontrada ---")
	fmt.Printf(MsgTarefaDetalhe+"\n", tarefa.ID, tarefa.Status, tarefa.Descricao)
}

func lerID(scanner *bufio.Scanner) (tarefas.TarefaID, error) {
	fmt.Print("Digite o ID da tarefa: ")
	scanner.Scan()
	idStr := scanner.Text()

	if idStr == "" {
		return 0, fmt.Errorf("ID não pode ser vazio")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter ID para número: %v", err)
	}

	if id <= 0 {
		return 0, fmt.Errorf("ID deve ser positivo")
	}

	return tarefas.TarefaID(id), nil
}
