package main

import (
	"fmt"
	"math"
)

// Função para calcular o volume da piscina
func calcularVolume(formato int) float64 {
	var volume float64
	switch formato {

	case 1: // Retangular
		var comprimento, largura, profundidade float64
		fmt.Print("Comprimento (m): ")
		fmt.Scan(&comprimento)
		fmt.Print("Largura (m): ")
		fmt.Scan(&largura)
		fmt.Print("Profundidade (m): ")
		fmt.Scan(&profundidade)
		if comprimento <= 0 || largura <= 0 || profundidade <= 0 {
			fmt.Println("Erro: Valores não podem ser negativos")
			return 0
		}
		volume = comprimento * largura * profundidade
	case 2: // Oval
		var diametroMaior, diametroMenor, profundidade float64
		fmt.Print("Diâmetro Maior (m): ")
		fmt.Scan(&diametroMaior)
		fmt.Print("Diâmetro Menor (m): ")
		fmt.Scan(&diametroMenor)
		fmt.Print("Profundidade (m): ")
		fmt.Scan(&profundidade)
		if diametroMaior <= 0 || diametroMenor <= 0 || profundidade <= 0 {
			fmt.Println("Erro: Valores não podem ser negativos")
			return 0
		}
		volume = (diametroMaior * diametroMenor * profundidade) * 0.785
	case 3: // Redonda
		var diametro, profundidade float64
		fmt.Print("Diâmetro (m): ")
		fmt.Scan(&diametro)
		fmt.Print("Profundidade (m): ")
		fmt.Scan(&profundidade)
		if diametro <= 0 || profundidade <= 0 {
			fmt.Println("Erro: Valores não podem ser negativos")
			return 0
		}
		volume = math.Pi * math.Pow(diametro/2, 2) * profundidade
	default:
		fmt.Println("Opção inválida")
		return 0
	}
	return volume
}

// Função para processar as opções de tratamento
func tratarPiscina(volume float64) {
	for {
		fmt.Println("\nEscolha uma opção de tratamento:")
		fmt.Println("1 - Dosagens Diárias e Semanais")
		fmt.Println("2 - Decantação para Aspirar")
		fmt.Println("3 - Correção de pH")
		fmt.Println("4 - Superdosagens")
		fmt.Println("5 - (Implementação futura)")
		fmt.Println("6 - FIM")
		fmt.Print("Opção: ")
		var opcao int
		fmt.Scan(&opcao)
		switch opcao {
		case 1:
			fmt.Printf("Dosagem Inicial Algistático: %.2fg\n", volume*14)
			fmt.Printf("Dosagem Semanal Algistático: %.2fg\n", volume*6)
			fmt.Printf("Dosagem Cloro Líquido: %.2fg\n", volume*25)
		case 2:
			fmt.Printf("Sulfato de Alumínio: %.2fg\n", volume*60)
			fmt.Printf("Barrilha Leve: %.2fg\n", volume*30)
		case 3:
			var phAtual float64
			fmt.Print("Informe o pH atual: ")
			fmt.Scan(&phAtual)
			if phAtual < 7.2 {
				fmt.Println("Recomenda-se adicionar barrilha para aumentar o pH.")
			} else if phAtual > 7.8 {
				fmt.Println("Recomenda-se adicionar redutor de pH para diminuir.")
			} else {
				fmt.Println("O pH está ideal.")
			}
		case 4:
			fmt.Printf("Superdosagem Algicida: %.2fg\n", volume*16)
			fmt.Printf("Superdosagem Cloro Líquido: %.2fg\n", volume*50)
		case 5:
			fmt.Println("Opção ainda não implementada.")
		case 6:
			fmt.Println("Encerrando o programa.")
			return
		default:
			fmt.Println("Opção inválida - escolha de 1 a 6")
		}
	}
}

func main() {
	fmt.Println("Escolha o formato da piscina:")
	fmt.Println("1 - Retangular")
	fmt.Println("2 - Oval")
	fmt.Println("3 - Redonda")
	fmt.Print("Opção: ")
	var formato int
	fmt.Scan(&formato)
	volume := calcularVolume(formato)
	if volume > 0 {
		fmt.Printf("O volume da piscina é: %.2f m³\n", volume)
		tratarPiscina(volume)
	}
}
