# REQUISITOS:FERRAMENTAPARACALCULAADATA

# PREVISTADEPARTOANIMAL

## Objetivo

Desenvolver um programaem Go que calcula a data prevista de parto e o
próximocicloestraldeum animal, com basenaraçaenadatadecobertura
fornecidaspelousuário.

## FuncionalidadesPrincipais

1. **CalcularaDatadeParto** :

- Com base na raça do animal, determinar o número de dias de
  gestação.
- Adicionaressesdiasàdatadecoberturaparaobteradataprevista
  departo.
2. **CalcularoPróximoCicloEstral** :
- Adicionarumintervalofixode 21 dias àdatadoúltimocicloestral
  paraobteradatadopróximociclo.
3. **VerificarAnoBissexto** :
- Determinarseumanoébissexto.
4. **InteraçãocomoUsuário** :
- Solicitaraousuárioadatadecoberturaearaçadoanimal.
- Oferecera opçãode calcular opróximo ciclo estral ou encerrar o
  programa.

## EstruturasdeDados

1. **Data** :

- Representaumadatacomdia,mêseano.
- Métodos:adicionarDias,calcularDataParto,calcularCicloEstral,ehBissexto.
2. **Animal** :
- Representaumanimalcomaraça.
- Métodos:obterDiasGestacao.


## RequisitosFuncionais

1. **MétodoadicionarDias** :

- **Entrada** :Númerodediasaseremadicionados.
- **Saída** :Novadataapósaadiçãodosdias.
- **Regras** :
    - Seodiaexceder 30 ,incrementaromês.
    - Seomêsexceder 12 ,incrementaroanoeajustaromêspara
      janeiro.
2. **MétodocalcularDataParto** :
- **Entrada** :Númerodediasdegestação.
- **Saída** :Dataprevistadeparto.
- **Regras** :UtilizarométodoadicionarDias.
3. **MétodocalcularCicloEstral** :
- **Entrada** :Númerodediasdocicloestral.
- **Saída** :Datadopróximocicloestral.
- **Regras** :UtilizarométodoadicionarDias.
4. **MétodoehBissexto** :
- **Entrada** :Ano.
- **Saída** :Booleanoindicandoseoanoébissexto.
- **Regras** :
- Anobissextosedivisívelpor 4 enãopor 100 ,oudivisívelpor
400.
5. **MétodoobterDiasGestacao** :
- **Entrada** :Raçadoanimal.
- **Saída** :Númerodediasdegestação.
- **Regras** :
- ZEBU: 290 dias.
- EUROFEU: 280 dias.
- CRUZADO:Médiade 290 e 280 dias.
- Raçainválida: 285 dias(padrão).
6. **Funçãomain** :
- **Fluxo** :
- Obterdatadecobertura.
- Obterraçadoanimal.
- Calculardatadeparto.
- Exibirdatadeparto.
- Ofereceropçãoparacalcularopróximocicloestralouencerrar.
7. **FunçãoobterDataCobertura** :
- **Entrada** :Datadecoberturafornecidapelousuário.
- **Saída** :EstruturaDatacomdia,mêseano.
8. **FunçãoobterRacaAnimal** :


- **Entrada** :Raçadoanimalfornecidapelousuário.
- **Saída** :EstruturaAnimalcomaraça.
9. **FunçãoexibirDataParto** :
- **Entrada** :Datadeparto.
- **Saída** :ExibiradatadepartonoformatoDD/MM/YYYY.
10. **FunçãocalcularProximoCicloEstral** :
- **Entrada** :Datadoúltimocicloestralfornecidapelousuário.
- **Saída** :Datadopróximocicloestral.
- **Regras** :Adicionar 21 diasàdatadoúltimocicloestral.

## RequisitosNãoFuncionais

1. **Usabilidade** :

- Interfacedelinhadecomandointuitivaefácildeusar.
- Mensagensclarasparaousuário.
2. **Manutenibilidade** :
- Códigobemestruturadoecomentado.
- Usodenomesdescritivosparavariáveisefunções.
3. **Desempenho** :
- Tempoderespostarápidoparacálculoseinteraçõescomousuário.
4. **Confiabilidade** :
- Tratamentodeentradasinválidas(e.g.,datasforadointervaloválido).
- Garantirqueadatainicialsejasempreválida.

## ExemplosdeUso

1. **CalcularDatadeParto** :

- Entrada:Datadecobertura= 01 / 01 / 2023 ,Raça=ZEBU.
- Saída:Datadeparto= 20 / 10 / 2023.
2. **CalcularPróximoCicloEstral** :
- Entrada:Datadoúltimocicloestral= 01 / 01 / 2023.
- Saída:Próximocicloestral= 22 / 01 / 2023.

## ConsideraçõesFinais

- O programa deve ser testado com diferentes entradas para garantir a
  precisãodoscálculos.
- Documentação adicional pode ser necessária para detalhar o
  funcionamentodecadamétodoefunção.


## Funcionalidade:CalcularDatadePartoePróximoCicloEstral

ComoumusuárioQuerocalcularadataprevistadepartoeopróximocicloestral
deumanimalParaplanejarmelhoromanejoreprodutivo

- Cenário:FluxoPrincipal-CalcularDatadeParto

```
Dado queousuárioinsereadatadecobertura“ 01 / 01 / 2023 ” E araçado
animal é “ZEBU” Quando o usuário solicita o cálculo da data de parto
Então adataprevistadepartodeveser“ 20 / 10 / 2023 ” E amensagem“Data
previstaparaoparto: 20 / 10 / 2023 ”deveserexibida
```
- Cenário:Exceção-RaçaInválida

```
Dado queousuárioinsereadatadecobertura“ 01 / 01 / 2023 ” E araçado
animalé“INVALIDA” Quando ousuáriosolicitaocálculodadatadeparto
Então amensagem“Raçainválida!Usandovalorpadrãode 285 dias.”deve
serexibida E adataprevistadepartodevesercalculadacom 285 diasde
gestação
```
- Cenário:Exceção-DatadeCoberturaInválida

```
Dado queousuárioinsereadatadecobertura“ 31 / 02 / 2023 ” E araçado
animal é “ZEBU” Quando o usuário solicita o cálculo da data de parto
Então amensagem“Datadecoberturainválida.Porfavor,insiraumadata
válida.”deveserexibida E ocálculodadatadepartonãodeveserrealizado
```
- Cenário:Borda-DatadeCoberturanoFinaldoAno

```
Dado queousuárioinsereadatadecobertura“ 30 / 12 / 2023 ” E araçado
animal é “ZEBU” Quando o usuário solicita o cálculo da data de parto
Então adataprevistadepartodeveser“ 29 / 09 / 2024 ” E amensagem“Data
previstaparaoparto: 29 / 09 / 2024 ”deveserexibida
```
- Cenário:FluxoPrincipal-CalcularPróximoCicloEstral

```
Dado que o usuário insere a data do último ciclo estral “ 01 / 01 / 2023 ”
Quando ousuáriosolicitaocálculodopróximocicloestral Então adatado
próximociclo estral deveser“ 22 / 01 / 2023 ” E a mensagem“Próximociclo
estralprevisto: 22 / 01 / 2023 ”deveserexibida
```
- Cenário:Exceção-DatadoCicloEstralInválida

```
Dado que o usuário insere a data do último ciclo estral “ 31 / 02 / 2023 ”
Quando o usuário solicita o cálculo do próximo ciclo estral Então a
mensagem“Datadocicloestralinválida.Porfavor,insiraumadataválida.”
deveserexibida E ocálculodopróximocicloestralnãodeveserrealizado
```

