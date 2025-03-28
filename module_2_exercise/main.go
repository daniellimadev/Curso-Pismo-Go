package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Estrutura para representar uma data com dia, mês e ano.
type Data struct {
	Dia int
	Mes int
	Ano int
}

// Estrutura para representar um animal com sua raça.
type Animal struct {
	Raca string
}

// Método para verificar se um ano é bissexto.
func (d Data) ehBissexto() bool {
	ano := d.Ano
	// Um ano é bissexto se for divisível por 4, exceto aqueles divisíveis por 100,
	// a menos que também sejam divisíveis por 400.
	return (ano%4 == 0 && ano%100 != 0) || ano%400 == 0
}

// Método para verificar se uma data é válida (dia, mês e ano dentro dos limites).
func (d Data) ehValida() bool {
	// Verifica se o mês está dentro do intervalo válido (1 a 12).
	if d.Dia < 1 || d.Mes < 1 || d.Mes > 12 {
		return false // Dia ou mês fora do intervalo válido
	}
	// Número de dias em cada mês (janeiro a dezembro).
	diasEmMês := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Ajusta os dias de fevereiro se o ano for bissexto
	if d.Mes == 2 && d.ehBissexto() {
		diasEmMês[1] = 29
	}

	return d.Dia <= diasEmMês[d.Mes-1] // Verifica se o dia está dentro do número de dias do mês
}

// Método para adicionar um número de dias a uma data.
func (d Data) adicionarDias(dias int) Data {
	novaData := d // Cria uma cópia da data original para não modificá-la diretamente
	for i := 0; i < dias; i++ {
		novaData.Dia++
		// Se o dia exceder o número de dias no mês, ajusta para o primeiro dia do próximo mês.
		if novaData.Dia > novaData.diasNoMes() {
			novaData.Dia = 1
			novaData.Mes++
			// Se o mês exceder 12, ajusta para janeiro do próximo ano.
			if novaData.Mes > 12 {
				novaData.Mes = 1
				novaData.Ano++
			}
		}
	}
	return novaData // Retorna a nova data calculada
}

// Método para obter o número de dias em um determinado mês de um ano.
func (d Data) diasNoMes() int {
	// Número de dias em cada mês (janeiro a dezembro).
	diasEmMes := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Ajusta os dias de fevereiro se o ano for bissexto.
	if d.Mes == 2 && d.ehBissexto() {
		return 29
	}
	return diasEmMes[d.Mes-1]
}

// Método para calcular a data prevista do parto, adicionando os dias de gestação à data de cobertura.
func (d Data) calcularDataParto(diasGestacao int) Data {
	return d.adicionarDias(diasGestacao)
}

// Método para calcular a data do próximo ciclo estral, adicionando 21 dias à data do último ciclo.
func (d Data) calcularCicloEstral() Data {
	return d.adicionarDias(21)
}

// Método para obter a quantidade de dias de gestação com base na raça do animal.
func (a Animal) obterDiasGestacao() int {
	switch strings.ToUpper(a.Raca) {
	case "ZEBU":
		return 292 // Número de dias de gestação para Zebu
	case "EUROPEU":
		return 280 // Número de dias de gestação para Europeu
	case "CRUZADO":
		return (292 + 280) / 2 // Média dos dias de gestação para animais cruzados
	default:
		fmt.Println("Raça inválida! Usando valor padrão de 285 dias.")
		return 285 // Valor padrão para raças não reconhecidas
	}
}

// Método para formatar a data no formato DD/MM/AAAA.
func (d Data) formatarData() string {
	return fmt.Sprintf("%02d/%02d/%04d", d.Dia, d.Mes, d.Ano)
}

// Função para obter uma data do usuário, validando o formato e a validade da data.
func obterData() (Data, bool) {
	var input string
	fmt.Scanln(&input) // Lê a entrada do usuário

	// Divide a string de entrada em dia, mês e ano.
	partes := strings.Split(input, "/")

	// Verifica se o formato da data é válido (DD/MM/AAAA).
	if len(partes) != 3 {
		fmt.Println("Formato inválido. Use DD/MM/YYYY.")
		return Data{}, false // Retorna uma data vazia e false se o formato for inválido
	}

	// Converte as partes da data de string para inteiros.
	diaStr, mêsStr, anoStr := partes[0], partes[1], partes[2]
	dia, err1 := strconv.Atoi(diaStr)
	mes, err2 := strconv.Atoi(mêsStr)
	ano, err3 := strconv.Atoi(anoStr)

	// Verifica se a conversão foi bem-sucedida.
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Valores inválidos. Use números para dia, mês e ano.")
		return Data{}, false // Retorna uma data vazia e false se os valores não forem números
	}

	// Cria uma estrutura Data com os valores convertidos.
	date := Data{Dia: dia, Mes: mes, Ano: ano}

	// Verifica se a data é válida (dentro dos limites de dia e mês).
	if !date.ehValida() {
		fmt.Println("Data inválida. Por favor, insira uma data válida.")
		return Data{}, false // Retorna uma data vazia e false se a data for inválida
	}

	return date, true // Retorna a data válida e true
}

// Função para obter a raça do animal do usuário.
func obterRacaAnimal() Animal {
	var entrada string
	fmt.Scanln(&entrada)
	return Animal{Raca: strings.TrimSpace(entrada)} // Remove espaços em branco da entrada
}

// Função para exibir a data prevista do parto formatada.
func exibirDataParto(dataParto Data) {
	fmt.Printf("Data prevista para o parto: %s\n", dataParto.formatarData())
}

// Função para calcular e exibir a data do próximo ciclo estral.
func calcularProximoCicloEstral() {
	fmt.Print("Digite a data do último ciclo estral (DD/MM/YYYY): ")
	ultimoCiclo, valid := obterData()
	if !valid {
		return // Se a data for inválida, não prossegue
	}

	nextCycle := ultimoCiclo.calcularCicloEstral()
	fmt.Printf("Próximo ciclo estral previsto: %s\n", nextCycle.formatarData())
}

func main() {
	fmt.Println("=== FERRAMENTA PARA CÁLCULO DE DATA PREVISTA DE PARTO ANIMAL ===")
	fmt.Print("Digite a data de cobertura (DD/MM/YYYY): ")
	dataDeCobertura, valid := obterData()
	if !valid {
		return // Se a data de cobertura for inválida, encerra
	}

	fmt.Print("Digite a raça do animal (ZEBU, EUROPEU, CRUZADO ou outra): ")
	animal := obterRacaAnimal()

	// Obtém o número de dias de gestação com base na raça do animal.
	diasDeGestacao := animal.obterDiasGestacao()

	// Calcula a data prevista do parto adicionando os dias de gestação à data de cobertura
	dataDeNascimento := dataDeCobertura.calcularDataParto(diasDeGestacao)
	exibirDataParto(dataDeNascimento)

	var resposta string
	fmt.Print("\nDeseja calcular o próximo ciclo estral? (S/N): ")
	fmt.Scanln(&resposta)
	resposta = strings.TrimSpace(strings.ToUpper(resposta))

	if resposta == "S" {
		calcularProximoCicloEstral()
	} else {
		fmt.Println("Programa encerrado.")
	}
}
