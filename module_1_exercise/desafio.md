# FERRAMENTA PARA TRATADORES DE PISCINA

## Parte 1: Entrada e Cálculo das Dimensões da Piscina

Nesta parte, o código lida com a entrada das dimensões da piscina com
base no formato escolhido (retangular, oval ou redonda) e calcula o volume
da piscina.

**Requisitos:**

1. O código deve começar pedindo ao usuário que escolha o formato da
   piscina, com as opções numeradas de 1 a 3.
2. O código deve ler a escolha do usuário e, em seguida, solicitar as
   dimensões da piscina com base no formato escolhido.
3. Se o usuário escolher o formato retangular (opção 1), o código deve
   solicitar o comprimento, largura e profundidade da piscina. O cálculo do
   volume deve ser feito multiplicando as três dimensões: Volume =
   Comprimento x Largura x Profundidade.
4. Se o usuário escolher o formato oval (opção 2), o código deve solicitar o
   diâmetro maior, o diâmetro menor e a profundidade da piscina. O
   cálculo do volume deve ser feito multiplicando os dois diâmetros e a
   profundidade, e em seguida, multiplicando o resultado por 0,785:
   Volume = (DiâmetroMaior * DiâmetroMenor * Profundidade) * 0,785.
5. Se o usuário escolher o formato redondo (opção 3), o código deve
   solicitar o diâmetro da piscina e a profundidade. O cálculo do volume
   deve ser feito usando a fórmula para o volume de um cilindro: Volume =
   π * (Diâmetro / 2) * (Diâmetro / 2) * Profundidade.
6. O código deve armazenar o valor do volume calculado na variável “M”
   para uso posterior.


## Parte 2: Escolha de Opções para Tratamento da Piscina

Nesta parte, o código permite ao usuário escolher entre várias opções de
tratamento de piscina com base no volume calculado anteriormente.

**Requisitos:**

1. O código deve exibir um menu com opções numeradas de 1 a 6 para que
   o usuário escolha uma ação.
2. O código deve ler a escolha do usuário e executar a ação
   correspondente com base na escolha.
3. Cada opção (1 a 6) deve executar uma ação específica, conforme
   descrito a seguir:

```
 Opção 1: “Dosagens Diárias e Semanais”
```
```
 O código deve calcular as dosagens para algistático e cloro
líquido com base no valor de “M” e exibir os resultados. Por
exemplo, para o algistático: DosagemInicialAlgistático = M *
14, DosagemSemanalAlgistático = M * 6,
DosagemCloroLíquido = M * 25.
 Opção 2: “Decantação para Aspirar”
```
```
 O código deve calcular as dosagens de sulfato de alumínio e
barrilha leve com base no valor de “M” e exibir os
resultados, juntamente com as instruções. Por exemplo,
SulfatoAlumínioGramas = M * 60, BarrilhaLeveGramas = M
* 30.
 Opção 3: “Correção de pH”
```
```
 O código deve solicitar o pH atual da piscina e, com base
nesse valor, recomendar o pH ideal para a piscina. Você
pode definir faixas de pH e recomendar um valor de pH com
base na faixa em que o pH atual se encontra.
 Opção 4: “Superdosagens”
```
```
 O código deve calcular as dosagens para algicida e cloro
líquido com base no valor de “M” e exibir os resultados,
juntamente com as instruções. Por exemplo,
SuperdosagemAlgicida = M * 16,
SuperdosagemCloroLíquido = M * 50.
 Opção 5: Esta parte do código está em branco, sugerindo que
precisa ser implementada. Você pode adicionar uma
funcionalidade adicional a esta opção, se necessário.
```
```
 Opção 6: “FIM”
```

```
 O código deve encerrar a execução.
```
4. O código deve tratar opções inválidas e exibir uma mensagem de erro
   se o usuário inserir uma opção inválida.

## CENÁRIOS DE TESTE

Funcionalidade: Cálculo do volume da piscina e seleção de tratamentos

```
Cenário: Cálculo de volume para piscina retangular com dimensões
válidas
```
```
Dado que o usuário selecionou a opção 1 (Retangular)
Quando informa comprimento 5m, largura 3m e profundidade
1.5m
Então o sistema deve calcular o volume como 22.5 m³
```
```
Cenário: Tentativa de entrada com valor negativo para diâmetro
redondo
```
```
Dado que o usuário selecionou a opção 3 (Redonda)
Quando informa diâmetro -4m e profundidade 2m
Então o sistema deve exibir "Erro: Valores não podem ser
negativos"
```
```
Cenário: Seleção de opção de tratamento inválida
```
```
Dado que o volume M foi calculado como 30m³
Quando o usuário seleciona a opção 7
Então o sistema deve mostrar "Opção inválida - escolha de 1
a 6"
```
Funcionalidade: Tratamentos químicos da piscina

```
Cenário: Cálculo de dosagem semanal para piscina média
```
```
Dado M = 50m³ e opção 1 selecionada
Quando processa dosagens
Então deve retornar 700g algistático inicial
E 300g semanal de algistático
```

