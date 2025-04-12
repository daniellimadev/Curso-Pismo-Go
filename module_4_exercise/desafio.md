# Especificação de Requisitos para o Sistema de Gerenciamento de Pedidos de Restaurante

# Objetivo

Desenvolver um sistema de gerenciamento de pedidos para um restaurante, permitindo
adicionar, listar e marcar pedidos como entregues. O sistema deve ser implementado
em Go e incluir testes unitários para garantir a funcionalidade correta.

# Requisitos Funcionais

1. **Menu Principal**
    - O sistema deve exibir um menu com as seguintes opções:
        1. Adicionar um novo pedido
        2. Listar todos os pedidos pendentes
        3. Marcar um pedido como entregue
        4. Sair
2. **Adicionar Pedido**
    - O usuário deve poder adicionar um novo pedido fornecendo uma descrição.
    - Cada pedido deve receber um número único e sequencial.
    - O status inicial do pedido deve ser “pendente”.
3. **Listar Pedidos**
    - O sistema deve listar todos os pedidos, mostrando o número do pedido,
      descrição e status (pendente ou entregue).
4. **Marcar Pedido como Entregue**
    - O usuário deve poder marcar um pedido como entregue fornecendo o
      número do pedido.
    - O sistema deve atualizar o status do pedido para “entregue”.
5. **Simulação de Preparo**
    - Após adicionar um pedido, o sistema deve simular um tempo de preparo de
      5 segundos.
    - Após o tempo de preparo, o pedido deve ser automaticamente marcado
      como entregue.
6. **Encerramento do Programa**
    - O sistema deve permitir que o usuário encerre o programa selecionando a
      opção “Sair”.

# Requisitos Não Funcionais

1. **Desempenho**
    - O sistema deve responder rapidamente às interações do usuário, com
      tempos de resposta inferiores a 1 segundo para operações que não
      envolvem simulação de preparo.
2. **Usabilidade**


- O sistema deve ser intuitivo e fácil de usar, com mensagens claras e
  instruções para o usuário.
3. **Testabilidade**
- O sistema deve incluir testes unitários para verificar a funcionalidade de
  adicionar pedidos e marcar pedidos como entregues.

# Estrutura do Código

1. **restaurante.go**
    - Contém a função principal (main) que inicializa o gerenciador de pedidos e
      entra em um loop infinito para processar as opções do menu.
    - Funções auxiliares: exibirMenu, lerOpcao, executarOpcao, adicionarNovoPedido,
      marcarPedidoComoEntregue.
2. **pedido.go**
    - Define a estrutura Pedido e a função NovoPedido para criar novos pedidos.
3. **gerenciadorpedidos.go**
    - Define a estrutura GerenciadorPedidos e métodos para adicionar pedidos,
      listar pedidos, marcar pedidos como entregues, simular preparo e verificar o
      status de entrega.
4. **restaurante_test.go**
    - Contém testes unitários para verificar a funcionalidade de adicionar pedidos
      (TestAdicionarPedido) e marcar pedidos como entregues
      (TestMarcarComoEntregue).

# Critérios de Aceitação

- O sistema deve permitir adicionar, listar e marcar pedidos como entregues
  conforme descrito.
- O sistema deve simular o tempo de preparo e atualizar o status do pedido
  automaticamente.
- Todos os testes unitários devem passar sem erros.
- O sistema deve ser robusto e lidar com entradas inválidas do usuário, exibindo
  mensagens de erro apropriadas.

# Considerações Finais

- O desenvolvedor deve seguir as melhores práticas de programação em Go,
  incluindo o uso de convenções de nomenclatura e estrutura de código.
- O código deve ser bem documentado, com comentários explicando a
  funcionalidade de cada parte do sistema.


# Funcionalidade: Gerenciamento de Pedidos de Restaurante

Como usuário do sistema Quero gerenciar pedidos de restaurante Para garantir que os
pedidos sejam processados e entregues corretamente

- Cenário: Adicionar um novo pedido com sucesso

**Dado** que estou na tela principal do sistema **Quando** seleciono a opção “Adicionar
um novo pedido” **E** insiro a descrição “Hambúrguer com batatas fritas” **Então** o
sistema adiciona o pedido com status “pendente” **E** exibe a mensagem “Pedido
adicionado com sucesso” **E** o pedido recebe um número único e sequencial

- Cenário: Listar todos os pedidos pendentes

**Dado** que existem pedidos no sistema **Quando** seleciono a opção “Listar todos os
pedidos pendentes” **Então** o sistema exibe uma lista com todos os pedidos **E** cada
pedido mostra o número, descrição e status (pendente ou entregue)

- Cenário: Marcar um pedido como entregue

**Dado** que existem pedidos pendentes no sistema **Quando** seleciono a opção
“Marcar um pedido como entregue” **E** insiro o número do pedido “1” **Então** o
sistema atualiza o status do pedido para “entregue” **E** exibe a mensagem “Pedido
1 marcado como entregue”

- Cenário: Tentar marcar um pedido inexistente como entregue

**Dado** que existem pedidos no sistema **Quando** seleciono a opção “Marcar um
pedido como entregue” **E** insiro o número do pedido “999” **Então** o sistema exibe a
mensagem de erro “Pedido 999 não encontrado”

- Cenário: Adicionar um pedido sem descrição

**Dado** que estou na tela principal do sistema **Quando** seleciono a opção “Adicionar
um novo pedido” **E** deixo o campo de descrição vazio **Então** o sistema exibe a
mensagem de erro “Descrição do pedido é obrigatória”

- Cenário: Simulação de preparo automático

**Dado** que adicionei um novo pedido **Então** o sistema simula um tempo de preparo
de 5 segundos **E** automaticamente atualiza o status do pedido para “entregue” **E**
exibe a mensagem “Pedido X entregue!”

- Cenário: Encerrar o programa

**Dado** que estou na tela principal do sistema **Quando** seleciono a opção “Sair”
**Então** o sistema encerra a execução **E** exibe a mensagem “Saindo...”



