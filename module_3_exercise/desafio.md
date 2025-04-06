# EspecificaçãodeRequisitos:Gerenciadorde

# TarefasCLI(Pacotetarefas)

**1 .VisãoGeral**

EstedocumentodescreveosrequisitosparaummóduloGo(packagetarefas)que
implementaumgerenciadordetarefassimplesoperadovialinhadecomando
(CLI).Osistemapermiteaosusuáriosadicionar,visualizar,marcarcomo
concluídaseremovertarefas.Astarefassãoarmazenadasemmemóriadurante
aexecuçãodoprogramaenãopersistemapósotérmino.Osistemadeveser
projetadoparausoporumúnicousuárioemumambientenãoconcorrente.

**2 .Escopo**

Oescopodestemóduloinclui:

- Definiçãodasestruturasdedadospararepresentartarefas.
- Implementaçãodalógicadenegócioparagerenciartarefas(adicionar,
  concluir,remover,listar,buscar).
- Definiçãodeumainterfaceparaalógicadenegócio,promovendo
  desacoplamento.
- ImplementaçãodefunçõesparainteraçãocomousuárioviaCLI(menu,
  leituradeentrada,exibiçãodemensagens).
- Definiçãodeconstantesparastatus,mensagensdeerroemensagensde
  UI.

**NÃOestánoescopo:**

- Persistênciadedadosemarquivosoubancodedados.
- Suporteamúltiplosusuáriossimultâneos(concorrência).
- Interfacegráfica(GUI)ouAPIweb.
- Autenticaçãoouautorizaçãodeusuários.

**3 .RequisitosFuncionais(RF)**

**RF 01 :AdicionarTarefa**

- **Descrição:** Osistemadevepermitiraousuárioadicionarumanovatarefa.
- **Entrada:** Umastringcontendoadescriçãodatarefa.
- **Processamento:**
    - Adescriçãofornecidadeveservalidada:nãopodeservaziaapós
      removerespaçosembrancodoinícioedofim(strings.TrimSpace).Se
      forinválida,umerro(ErroDescricaoVazia)deveserretornado.
    - UmanovatarefadevesercriadacomumIDúnicoesequencial,
      começandoem 1.
    - OstatusinicialdanovatarefadeveserStatusPendente.


- Atarefacriadadeveserarmazenadanacoleçãointernadetarefas.
- OpróximoIDaserutilizadodeveserincrementado.
- **Saída:**
- Emcasodesucesso:ATarefarecém-criadaenilcomoerro.NaUI,
  exibiramensagemMsgTarefaAdicionadaformatadacomoIDdatarefa.
- Emcasodefalha(descriçãovazia):UmaTarefavaziaeoerro
  ErroDescricaoVazia.NaUI,exibiramensagemMsgErroEntradaDesce
  detalhesdoerro.

**RF 02 :MarcarTarefacomoConcluída**

- **Descrição:** Osistemadevepermitiraousuáriomarcarumatarefaexistente
  comoconcluída.
- **Entrada:** OTarefaIDdatarefaasermarcada.
- **Processamento:**
    - OsistemadevebuscaratarefapeloIDfornecido.
    - Seatarefanãoforencontrada,umerro(ErroTarefaNaoEncontrada)
      deveserretornado.
    - SeatarefaforencontradaeseustatusatualforStatusPendente,o
      statusdeveseralteradoparaStatusConcluida.
    - SeatarefaforencontradaeseustatusjáforStatusConcluida,
      nenhumaalteraçãodeveserfeitaeaoperaçãodeveserconsiderada
      bem-sucedida(idempotente).
- **Saída:**
    - Emcasodesucesso(tarefaencontradaemarcadaoujáconcluída):
      nilcomoerro.NaUI,exibiramensagemMsgTarefaConcluida
      formatadacomoIDdatarefa.
    - Emcasodefalha(tarefanãoencontrada):Oerro
      ErroTarefaNaoEncontrada.NaUI,exibiramensagem
      MsgErroMarcarConcluirformatadacomoerro.

**RF 03 :RemoverTarefa**

- **Descrição:** Osistemadevepermitiraousuárioremoverumatarefa
  existente.
- **Entrada:** OTarefaIDdatarefaaserremovida.
- **Processamento:**
    - OsistemadevebuscaratarefapeloIDfornecido.
    - Seatarefanãoforencontrada,umerro(ErroTarefaNaoEncontrada)
      deveserretornado.
    - Seatarefaforencontrada,eladeveserremovidapermanentemente
      dacoleçãointernadetarefas.
- **Saída:**


- Emcasodesucesso(tarefaencontradaeremovida):nilcomoerro.
  NaUI,exibiramensagemMsgTarefaRemovidaformatadacomoIDda
  tarefa.
- Emcasodefalha(tarefanãoencontrada):Oerro
  ErroTarefaNaoEncontrada.NaUI,exibiramensagemMsgErroRemover
  formatadacomoerro.

**RF 04 :ListarTarefasPendentes**

- **Descrição:** Osistemadevelistartodasastarefascomstatus
  StatusPendente.
- **Entrada:** Nenhumaentradaexplícitadousuário(aopçãodemenuimplicao
  filtro).
- **Processamento:**
    - Osistemadeverecuperartodasastarefasdacoleçãointernacujo
      statussejaStatusPendente.
    - AlistaresultantedeveserordenadacrescentementepeloTarefaID.
- **Saída:**
    - Umaslicecontendoastarefaspendentesordenadasenilcomoerro.
    - NaUI:
        - ExibirotítuloMsgListandoTarefasformatadocom“Pendentes”.
        - Sealistaestivervazia,exibirMsgNenhumaTarefa.
        - Paracadatarefanalista,exibirseusdetalhesusandoo
          formatoMsgTarefaDetalhe(ID,Status,Descrição).

**RF 05 :ListarTodasasTarefas**

- **Descrição:** Osistemadevelistartodasastarefas,independentementedo
  status.
- **Entrada:** Nenhumaentradaexplícitadousuário(aopçãodemenuimplica
  semfiltro).
- **Processamento:**
    - Osistemadeverecuperartodasastarefasdacoleçãointerna.
    - AlistaresultantedeveserordenadacrescentementepeloTarefaID.
- **Saída:**
    - Umaslicecontendotodasastarefasordenadasenilcomoerro.
    - NaUI:
        - ExibirotítuloMsgListandoTarefasformatadocom“Todas”.
        - Sealistaestivervazia,exibirMsgNenhumaTarefa.
        - Paracadatarefanalista,exibirseusdetalhesusandoo
          formatoMsgTarefaDetalhe(ID,Status,Descrição).

**RF 06 :BuscarTarefaporID**

- **Descrição:** Osistemadevepermitiraousuáriobuscareexibirosdetalhes
  deumatarefaespecíficapeloseuID.


- **Entrada:** OTarefaIDdatarefaaserbuscada.
- **Processamento:**
    - OsistemadevebuscaratarefapeloIDfornecidonacoleçãointerna.
    - Seatarefanãoforencontrada,umerro(ErroTarefaNaoEncontrada)
      deveserretornado.
- **Saída:**
    - Emcasodesucesso(tarefaencontrada):ATarefaencontradaenil
      comoerro.NaUI,exibirumtítulo“—TarefaEncontrada—”eos
      detalhesdatarefausandoMsgTarefaDetalhe.
    - Emcasodefalha(tarefanãoencontrada):UmaTarefavaziaeoerro
      ErroTarefaNaoEncontrada.NaUI,exibiramensagemMsgErroBuscar
      formatadacomoerro.

**RF 07 :SairdoPrograma**

- **Descrição:** Osistemadevepermitiraousuárioencerraraaplicação.
- **Entrada:** Aopçãodemenucorrespondentea“Sair”.
- **Processamento:** Oprogramadeveterminarsuaexecuçãodeformalimpa.
- **Saída:** NaUI,exibiramensagemMsgSaindo.

**RF 08 :ValidaçãodeEntradadeID**

- **Descrição:** OsistemadevevalidaraentradadousuárioquandoumIDde
  tarefaforsolicitado.
- **Entrada:** Umastringfornecidapelousuário.
- **Processamento:**
    - Astringnãopodeservazia.
    - Astringdeveserconversívelparaumnúmerointeiro(strconv.Atoi).
    - Onúmerointeiroresultantedeveserpositivo(> 0 ).
- **Saída:**
    - Emcasodesucesso:OTarefaIDválido.
    - Emcasodefalha:Umerroindicandoacausa(vazio,nãonumérico,
      nãopositivo).NaUI,exibiramensagemMsgErroEntradaIDedetalhes
      doerro.

**RF 09 :TratamentodeOpçãodeMenuInválida**

- **Descrição:** Seousuárioinserirumaopçãodemenuquenãocorrespondea
  nenhumaaçãoválida.
- **Entrada:** Umastringounúmerocorrespondenteàescolhadomenu.
- **Processamento:** Osistemadeveidentificarqueaopçãonãoestánalista
  deopçõesválidas[ 1 - 7 ].
- **Saída:** NaUI,exibiramensagemMsgOpcaoInvalidaereexibiromenu
  principal(MenuPrompt).

**RF 10 :LeituradeEntradadeTextoGenérica**


- **Descrição:** Osistemadeveterumaformapadronizadadelertextodo
  usuário.
- **Entrada:** Umastringdepromptaserexibida(opcional).
- **Processamento:** Lerumalinhadetextodaentradapadrão.Remover
  espaçosembrancodoinícioefim.Tratarpossíveiserrosdeleitura
  (fmt.Scanln),incluindoacondiçãounexpectednewlinequedevesertratada
  comoentradavaziasemerroexplícitonaleitura,maspodeservalidada
  posteriormente(e.g.,RF 01 ).
- **Saída:** Astringlidaeprocessada,ouumerrosealeiturafalhar(exceto
  unexpectednewline).

**4 .RequisitosNãoFuncionais(RNF)**

- **RNF 01 :PersistênciadeDados:** Osdadosdastarefasdevemsermantidos
  **apenasemmemória** .Aoencerraroprograma,todososdadosserão
  perdidos.
- **RNF 02 :Concorrência:** Osistema **nãoéthread-safe** .Nãodevehaver
  garantiasdecomportamentocorretoseasfunçõesdogerenciadorde
  tarefasforemchamadasconcorrentementepormúltiplasgoroutines.A
  implementação _não_ deveusarmutexesououtrosmecanismosde
  sincronização.
- **RNF 03 :Usabilidade:** Ainterfacedeveservialinhadecomando(CLI),
  apresentandoummenuclarocomopçõesnumeradas.Asmensagenspara
  ousuário(prompts,sucesso,erro)devemserinformativaseconsistentes,
  conformedefinidonasconstantes.Umapausaopcional(Pausar)podeser
  implementadaparamelhoraralegibilidadedasaídanoterminal.
- **RNF 04 :Desempenho:** Asoperaçõesdebusca,adição,marcaçãoe
  remoçãoporIDdevemtercomplexidademédiadeO( 1 )(característicade
  mapas).Alistagem(comousemfiltro)terácomplexidadedeO(N)para
  iteraçãoeO(NlogN)paraordenação,ondeNéonúmerototaldetarefas.
  Estedesempenhoéaceitávelparaoescopodefinido(gerenciamentode
  pequenoamédiovolumedetarefasemmemória).
- **RNF 05 :Manutenibilidade:**
    - OcódigodeveserescritoemGo,seguindoasconvençõespadrãoda
      linguagem(gofmt).
    - Devehaverclaraseparaçãoderesponsabilidades:
        - Estruturasdedados(Tarefa,TarefaID,StatusTarefa).
        - InterfacedeLógicadeNegócio(GerenciadorTarefas).
        - ImplementaçãodaLógicadeNegócio(ListaDeTarefas).
        - FunçõesdeUI/Interação(ExibirMenu,LerEntradaTexto,
          LerEntradaID,Handlers).
    - Utilizarconstantesnomeadasparastatus,mensagensdeerroe
      stringsdeUIparafacilitaramodificaçãoeinternacionalização
      (emboraestaúltimanãosejaumrequisito).


- AimplementaçãoListaDeTarefasdevesatisfazerainterface
  GerenciadorTarefas(verificávelcomvar_GerenciadorTarefas=
  (*ListaDeTarefas)(nil)).

**5 .ModelodeDados**

- **TarefaID** :Tipoint.Representaoidentificadorúnicodeumatarefa.Deveser
  sempreumvalorpositivo.
- **StatusTarefa** :Tipostring.Representaoestadodeumatarefa.Valores
  permitidos:
  - "pendente"(ConstanteStatusPendente)
  - "concluida"(ConstanteStatusConcluida)
- **Tarefa** :Estrutura(struct)contendoosseguintescampos:
    - IDTarefaID:Oidentificadorúnicodatarefa.
    - Descricaostring:Otextodescritivodatarefa.
    - StatusStatusTarefa:Oestadoatualdatarefa.
- **ArmazenamentoInterno:** Aimplementação(ListaDeTarefas)deveusarum
  map[TarefaID]Tarefaparaarmazenarastarefas,permitindoacessorápidopor
  ID.
- **GeraçãodeID:** Aimplementação(ListaDeTarefas)devemanterumcontador
  interno(proximoIDdotipoTarefaID)inicializadocom 1 ,queéusadopara
  atribuiroIDàpróximatarefaadicionadaedepoisincrementado.

**6 .InterfaceGerenciadorTarefas**

DeveserdefinidaumainterfaceGochamadaGerenciadorTarefascomos
seguintesmétodos,servindocomocontratoparaalógicadenegócio:

**type** GerenciadorTarefas **interface** {
AdicionarTarefa(descricaostring)(Tarefa,error)
MarcarComoConcluida(idTarefaID)error
RemoverTarefa(idTarefaID)error
ListarTarefas(filtroStatusStatusTarefa)([]Tarefa,error) _//filtroStatusvaziolistatodas_
BuscarTarefaPorID(idTarefaID)(Tarefa,error)
}

**7 .TratamentodeErros**

Osistemadeveusarerrosnomeadose/ouformatadosparaindicarcondições
específicasdefalha.Osseguinteserrosdenegóciodevemserdefinidoscomo
variáveisglobais(usandoerrors.New):

- ErroTarefaNaoEncontrada:Retornadoquandoumaoperaçãotentaacessar
  umatarefacomumIDquenãoexiste.
- ErroDescricaoVazia:RetornadoporAdicionarTarefaseadescriçãoforinválida.
- ErroStatusInvalido:RetornadoporListarTarefasseumfiltroStatusinválidofor
  fornecido(emboranaUIatual,apenasfiltrosválidosouvaziosejam
  passados).


- ErroTarefaJaConcluida:(Opcional,maspresentenocódigooriginal)Podeser
  retornadoporMarcarComoConcluidaseatarefajáestiverconcluída,masa
  especificaçãoatual(RF 02 )definecomportamentoidempotente(retornarnil).
  Manterocomportamentoidempotenteépreferível.

Errostécnicos/entradadevemsertratadose,senecessário,encapsulados
(fmt.Errorfcom%w):

- Erronaleituradaentradadousuário(fmt.Scanln).
- ErronaconversãodeIDstringparaint(strconv.Atoi).
- ErroparaIDvazioounãopositivo.

Mensagensdeerroespecíficasparaousuáriodevemserusadasnacamadade
UI(Handlers),conformedefinidonasconstantesMsgErro....

**8 .Constantes**

Todasasstringsliteraisusadasparastatus,prompts,mensagensde
sucesso/erroetítulosdevemserdefinidascomoconstantesnoiníciodoarquivo
parafácilmanutenção.(Videlistacompletanocódigooriginal).

**9 .FluxodaAplicação(CLI)**

1. InicializarumGerenciadorTarefas(instânciadeNovaListaDeTarefas).
2. Entraremumloopinfinito:

```
a. Exibiromenudeopções(ExibirMenu).
b. Leraescolhadousuário(LerEntradaTexto).
c. Converteraescolhaparanúmero.
d. Executaraaçãocorrespondenteusandoumswitchnaescolha:
*Chamarohandlerapropriado(HandleAdicionarTarefa,
HandleMarcarConcluida,etc.),passandoainstânciado
GerenciadorTarefas.
```
- Paraaopção“Sair”,exibirMsgSaindoeretornar/quebraroloop.
- Paraopçõesinválidas,exibirMsgOpcaoInvalida.

3. Fimdoprograma.

**10 .ConsideraçõesFinais**

Estaespecificaçãodetalhaocomportamentoesperadodopacotetarefas.O
desenvolvedordeveimplementarafuncionalidadeaderindoaestesrequisitos,
garantindoqueaestrutura,osnomesdetipos/funções/constantes,o
comportamentoeotratamentodeerroscorrespondamaodescrito.


## Funcionalidade:GerenciamentodeTarefasviaCLI

ComoumusuáriodaaplicaçãoCLI,Desejogerenciarminhastarefas(adicionar,
concluir,remover,listar,buscar)Paraorganizarmeutrabalhodeformaeficiente
atravésdalinhadecomando.

**Cenário** :FluxoPrincipal-AdicionareListarTodasasTarefasDadoqueo
gerenciadordetarefasestávazioQuandoousuárioescolheraopção“ 1 ”
(AdicionarTarefa)Einseriradescrição“Comprarleitenosupermercado”Então
osistemadeveexibiramensagem“Tarefa(ID: 1 )adicionadacomsucesso.”
Quandoousuárioescolheraopção“ 5 ”(ListarTodasasTarefas)Entãoo
sistemadeveexibirotítulo“—TarefasTodas—”Eosistemadeveexibiros
detalhesdatarefa“ID: 1 |Status:pendente|Descrição:Comprarleiteno
supermercado”#Verificaofluxobásicodeadicionarumatarefa(RF 01 )evê-la
nalistagemgeral(RF 05 ),#incluindoamensagemdesucesso
(MsgTarefaAdicionada)eoformatodelistagem(MsgTarefaDetalhe).

**Cenário** :Exceção-AdicionarTarefacomDescriçãoVaziaDadoqueo
gerenciadordetarefasestáiniciadoQuandoousuárioescolheraopção“ 1 ”
(AdicionarTarefa)Einserirumadescriçãocontendoapenasespaçosembranco
""Entãoosistemadeveexibiramensagem“Descriçãoinválida.”Eosistema
deveexibirodetalhedoerro“Detalhe:adescriçãodatarefanãopodeservazia”
Enenhumanovatarefadeveseradicionadaaogerenciador#Garantequea
validaçãodedescriçãovazia(RF 01 ,TrimSpace)funciona,#exibindoa
mensagemcorreta(MsgErroEntradaDesc)eoerroespecífico
(ErroDescricaoVazia).

**Cenário** :Exceção-MarcarTarefaInexistentecomoConcluídaDadoqueo
gerenciadordetarefascontématarefaID 1 :“Revisardocumento”comstatus
“pendente”Quandoousuárioescolheraopção“ 2 ”(MarcarcomoConcluída)E
inseriroID“ 999 ”(quenãoexistenalista)Entãoosistemadeveexibira
mensagem“Erroaomarcartarefacomoconcluída:tarefanãoencontrada”Ea
tarefaID 1 devepermanecercomstatus“pendente”#Testaotratamentodeerro
aotentaroperar(RF 02 )emumIDquenãoexiste,#verificandoamensagemde
erro(MsgErroMarcarConcluir)combinadacomoerro
(ErroTarefaNaoEncontrada).#Implicitamentecobrepartedavalidaçãode
entradadeID(RF 08 ).


**Cenário** :Borda-ListarTarefasPendentesQuandoNãoHáNenhumaDadoque
ogerenciadordetarefasestávazioQuandoousuárioescolheraopção“ 4 ”
(ListarTarefasPendentes)Entãoosistemadeveexibirotítulo“—Tarefas
Pendentes—”Eosistemadeveexibiramensagem“Nenhumatarefa
encontrada.”#Verificaocomportamentodalistagem(RF 04 )nocasolimitede
listavazia,#garantindoqueamensagemcorreta(MsgNenhumaTarefa)seja
exibida.

**Cenário** :Borda/ComportamentoEspecífico-TentarMarcarTarefaJáConcluída
DadoqueogerenciadordetarefascontématarefaID 1 :“Enviaremail
importante”comstatus“concluida”Quandoousuárioescolheraopção“ 2 ”
(MarcarcomoConcluída)EinseriroID“ 1 ”Entãoosistemadeveexibira
mensagem“Tarefa(ID: 1 )marcadacomoconcluída.”EostatusdatarefaID 1
devepermanecer“concluida”#Confirmaocomportamentoidempotentedefinido
noRF 02 .Mesmoqueatarefajáestejaconcluída,#aoperaçãoéconsiderada
sucessoeamensagem(MsgTarefaConcluida)éexibida,semerros.

**Cenário** :Exceção-FornecerIDInválido(NãoNumérico)paraBuscarTarefa
DadoqueogerenciadordetarefasestáiniciadoQuandoousuárioescolhera
opção“ 6 ”(BuscarTarefaporID)EinseriroID“abc”(nãonumérico)Entãoo
sistemadeveexibiramensagem“IDinválido.”Eosistemadeveexibirodetalhe
doerrocontendo“erroaoconverterIDparanúmero”#Testaavalidaçãode
entradadeID(RF 08 ),especificamenteparaformatonãonumérico(strconv.Atoi),
#verificandoamensagemgenérica(MsgErroEntradaID)eodetalhedoerro
técnico.

**Cenário** :Exceção-EscolherOpçãodeMenuInválidaDadoqueomenu
principaléexibidoQuandoousuárioescolheraopção“ 9 ”(quenãoestáno
menu[ 1 - 7 ])Entãoosistemadeveexibiramensagem“Opçãoinválida.Tente
novamente.”Eomenuprincipaldeveserexibidonovamente#Verificao
tratamentoparaentradasquenãocorrespondemaopçõesválidas(RF 09 ),#
garantindoaexibiçãodamensagem(MsgOpcaoInvalida)eareapresentaçãodo
menu.

- --


