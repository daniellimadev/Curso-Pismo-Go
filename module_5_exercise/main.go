package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	gerenciador := NovoGerenciador()

	http.HandleFunc("/compromissos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			gerenciador.CriarCompromisso(w, r)
		case http.MethodGet:
			gerenciador.ListarCompromissos(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/compromissos/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			http.Error(w, "URL inválida", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			gerenciador.ObterCompromisso(w, r, id)
		case http.MethodPut:
			gerenciador.AtualizarCompromisso(w, r, id)
		case http.MethodDelete:
			gerenciador.ExcluirCompromisso(w, r, id)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
