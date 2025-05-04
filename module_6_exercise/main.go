package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	gerenciador := NovoGerenciador()

	e.POST("/compromissos", func(c echo.Context) error {
		var compromisso Compromisso
		if err := c.Bind(&compromisso); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Requisição inválida"})
		}
		resp := make(chan interface{})
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Criar,
			Compromisso:  compromisso,
			RespostaChan: resp,
		}
		result := <-resp
		return c.JSON(http.StatusCreated, result)
	})

	e.GET("/compromissos", func(c echo.Context) error {
		resp := make(chan interface{})
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Listar,
			RespostaChan: resp,
		}
		result := <-resp
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/compromissos/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "ID inválido"})
		}
		resp := make(chan interface{})
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Buscar,
			ID:           id,
			RespostaChan: resp,
		}
		result := <-resp
		if err, ok := result.(error); ok {
			return c.JSON(http.StatusNotFound, map[string]string{"erro": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	})

	e.PUT("/compromissos/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "ID inválido"})
		}
		var compromisso Compromisso
		if err := c.Bind(&compromisso); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Requisição inválida"})
		}
		resp := make(chan interface{})
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Atualizar,
			ID:           id,
			Compromisso:  compromisso,
			RespostaChan: resp,
		}
		result := <-resp
		if err, ok := result.(error); ok {
			return c.JSON(http.StatusNotFound, map[string]string{"erro": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	})

	e.DELETE("/compromissos/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "ID inválido"})
		}
		resp := make(chan interface{})
		gerenciador.requisicoes <- Requisicao{
			Operacao:     Excluir,
			ID:           id,
			RespostaChan: resp,
		}
		result := <-resp
		if err, ok := result.(error); ok {
			return c.JSON(http.StatusNotFound, map[string]string{"erro": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
