package main

import (
	"encoding/json"
	"net/http"
)

// GerenciadorCompromissos gerencia os compromissos
type GerenciadorCompromissos struct {
	compromissos map[int]Compromisso
	proximoID    int
}

// NovoGerenciador cria um novo gerenciador de compromissos
func NovoGerenciador() *GerenciadorCompromissos {
	return &GerenciadorCompromissos{
		compromissos: make(map[int]Compromisso),
		proximoID:    1,
	}
}

// CriarCompromisso adiciona um novo compromisso
func (g *GerenciadorCompromissos) CriarCompromisso(w http.ResponseWriter, r *http.Request) {
	var compromisso Compromisso
	err := json.NewDecoder(r.Body).Decode(&compromisso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	compromisso.ID = g.proximoID
	g.compromissos[g.proximoID] = compromisso
	g.proximoID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(compromisso)
}

// ListarCompromissos retorna todos os compromissos
func (g *GerenciadorCompromissos) ListarCompromissos(w http.ResponseWriter, r *http.Request) {
	var lista []Compromisso
	for _, compromisso := range g.compromissos {
		lista = append(lista, compromisso)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lista)
}

// ObterCompromisso retorna um compromisso específico
func (g *GerenciadorCompromissos) ObterCompromisso(w http.ResponseWriter, r *http.Request, id int) {
	compromisso, existe := g.compromissos[id]
	if !existe {
		http.Error(w, "Compromisso não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(compromisso)
}

// AtualizarCompromisso atualiza um compromisso existente
func (g *GerenciadorCompromissos) AtualizarCompromisso(w http.ResponseWriter, r *http.Request, id int) {
	_, existe := g.compromissos[id]
	if !existe {
		http.Error(w, "Compromisso não encontrado", http.StatusNotFound)
		return
	}

	var compromisso Compromisso
	err := json.NewDecoder(r.Body).Decode(&compromisso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	compromisso.ID = id
	g.compromissos[id] = compromisso

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(compromisso)
}

// ExcluirCompromisso remove um compromisso
func (g *GerenciadorCompromissos) ExcluirCompromisso(w http.ResponseWriter, r *http.Request, id int) {
	_, existe := g.compromissos[id]
	if !existe {
		http.Error(w, "Compromisso não encontrado", http.StatusNotFound)
		return
	}

	delete(g.compromissos, id)
	w.WriteHeader(http.StatusNoContent)
}
