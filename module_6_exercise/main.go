package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inicializa uma nova instância do Echo
	e := echo.New()

	// Cria uma nova instância do Gerenciador de Compromissos
	gerenciador := NovoGerenciador()

	// Rota para criar um novo compromisso (POST /compromissos)
	e.POST("/compromissos", func(c echo.Context) error {
		var compromisso Compromisso

		// Faz o bind do corpo da requisição para a struct Compromisso
		if err := c.Bind(&compromisso); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Requisição inválida"})
		}

		// Cria canal de resposta e envia a requisição de criação ao gerenciador
		resp := make(chan interface{})
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Criar,
			Compromisso:  compromisso,
			RespostaChan: resp,
		}

		// Aguarda resposta e retorna com status 201 Created
		result := <-resp
		return c.JSON(http.StatusCreated, result)
	})

	// Rota para listar todos os compromissos (GET /compromissos)
	e.GET("/compromissos", func(c echo.Context) error {
		resp := make(chan interface{})

		// Envia requisição de listagem ao gerenciador
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Listar,
			RespostaChan: resp,
		}

		// Aguarda resposta e retorna com status 200 OK
		result := <-resp
		return c.JSON(http.StatusOK, result)
	})

	// Rota para buscar um compromisso por ID (GET /compromissos/:id)
	e.GET("/compromissos/:id", func(c echo.Context) error {
		// Converte o ID da URL para inteiro
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "ID inválido"})
		}

		resp := make(chan interface{})

		// Envia requisição de busca ao gerenciador
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Buscar,
			ID:           id,
			RespostaChan: resp,
		}

		// Aguarda resposta
		result := <-resp
		if err, ok := result.(error); ok {
			// Se não encontrar, retorna 404
			return c.JSON(http.StatusNotFound, map[string]string{"erro": err.Error()})
		}

		// Retorna o compromisso encontrado
		return c.JSON(http.StatusOK, result)
	})

	// Rota para atualizar um compromisso (PUT /compromissos/:id)
	e.PUT("/compromissos/:id", func(c echo.Context) error {
		// Converte o ID da URL para inteiro
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "ID inválido"})
		}

		var compromisso Compromisso

		// Faz o bind do corpo da requisição para a struct Compromisso
		if err := c.Bind(&compromisso); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Requisição inválida"})
		}

		resp := make(chan interface{})

		// Envia requisição de atualização ao gerenciador
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Atualizar,
			ID:           id,
			Compromisso:  compromisso,
			RespostaChan: resp,
		}

		// Aguarda resposta
		result := <-resp
		if err, ok := result.(error); ok {
			// Se não encontrar, retorna 404
			return c.JSON(http.StatusNotFound, map[string]string{"erro": err.Error()})
		}

		// Retorna o compromisso atualizado
		return c.JSON(http.StatusOK, result)
	})

	// Rota para excluir um compromisso (DELETE /compromissos/:id)
	e.DELETE("/compromissos/:id", func(c echo.Context) error {
		// Converte o ID da URL para inteiro
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "ID inválido"})
		}

		resp := make(chan interface{})

		// Envia requisição de exclusão ao gerenciador
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Excluir,
			ID:           id,
			RespostaChan: resp,
		}

		// Aguarda resposta
		result := <-resp
		if err, ok := result.(error); ok {
			// Se não encontrar, retorna 404
			return c.JSON(http.StatusNotFound, map[string]string{"erro": err.Error()})
		}

		// Retorna status 204 No Content se excluído com sucesso
		return c.NoContent(http.StatusNoContent)
	})

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}
