package main

// Compromisso representa um compromisso com título, descrição, data e hora.
// Os campos possuem tags JSON para que possam ser serializados e desserializados corretamente em requisições e respostas HTTP.
type Compromisso struct {
	ID          int    `json:"id"`          // Identificador único do compromisso
	Title       string `json:"title"`       // Título do compromisso
	Description string `json:"description"` // Descrição detalhada do compromisso
	Date        string `json:"date"`        // Data do compromisso no formato YYYY-MM-DD
	Time        string `json:"time"`        // Hora do compromisso no formato HH:MM
}
