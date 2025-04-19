**Especificações de Requisitos para o Sistema de Gerenciamento de
Compromissos**

**Objetivo** Desenvolver um sistema de gerenciamento de compromissos que
permita a criação, leitura, atualização e exclusão de compromissos. Cada
compromisso deve conter um título, uma descrição, uma data e um horário.
O sistema deve ser implementado em Go e seguir a estrutura e práticas de
codificação fornecidas nos exemplos de código.

**Requisitos Funcionais**

1. **Inicialização do Servidor**
    - O servidor deve ser inicializado na porta 8080.
    - O servidor deve imprimir a mensagem “Servidor rodando na porta
      8080” ao iniciar.
2. **Gerenciador de Compromissos**
    - O gerenciador de compromissos deve ser inicializado com um mapa
      vazio de compromissos, um ID inicial de 1 e um canal para solicitações
      de compromissos.
    - O gerenciador deve iniciar uma goroutine para processar solicitações
      de compromissos.
3. **Manipulação de Rotas**
    - O servidor deve definir as seguintes rotas:
      **-** /compromissos: Para criar e listar compromissos.
      **-** /compromissos/{id}: Para obter, atualizar e excluir um com-
      promisso específico.
4. **Criação de Compromissos**
    - A rota/compromissosdeve aceitar requisições POST para criar um
      novo compromisso.
    - O corpo da requisição deve conter um JSON com os campostitle,
      description,dateetime.
    - O servidor deve retornar o compromisso criado com um status HTTP
      201 Created.
5. **Listagem de Compromissos**
    - A rota/compromissosdeve aceitar requisições GET para listar todos
      os compromissos.
    - O servidor deve retornar uma lista de compromissos com um status
      HTTP 200 OK.
6. **Obtenção de Compromisso por ID**
    - A rota/compromissos/{id}deve aceitar requisições GET para obter
      um compromisso específico.
    - O servidor deve retornar o compromisso encontrado com um status
      HTTP 200 OK.
    - Se o compromisso não for encontrado, o servidor deve retornar um
      status HTTP 404 Not Found.
7. **Atualização de Compromisso**


- A rota/compromissos/{id}deve aceitar requisições PUT para atu-
  alizar um compromisso existente.
- O corpo da requisição deve conter um JSON com os campostitle,
  description,dateetime.
- O servidor deve retornar o compromisso atualizado com um status
  HTTP 200 OK.
- Se o compromisso não for encontrado, o servidor deve retornar um
  status HTTP 404 Not Found.
8. **Exclusão de Compromisso**
- A rota/compromissos/{id}deve aceitar requisições DELETE para
  excluir um compromisso existente.
- O servidor deve retornar um status HTTP 204 No Content.
- Se o compromisso não for encontrado, o servidor deve retornar um
  status HTTP 404 Not Found.

**Requisitos Não Funcionais**

1. **Estrutura do Projeto**
    - O projeto deve conter os seguintes arquivos:
      **-** main.go: Arquivo principal que inicializa o servidor e define os
      manipuladores de rotas.
      **-** compromisso.go: Arquivo que define a estrutura de um compro-
      misso.
      **-** gerenciador-compromisso.go: Arquivo que define o gerenci-
      ador de compromissos.
2. **Práticas de Codificação**
    - O código deve seguir as práticas de codificação e estrutura fornecidas
      nos exemplos de código.
    - Não devem ser utilizados pacotes ou bibliotecas adicionais além dos
      fornecidos nos exemplos.
3. **Tratamento de Erros**
    - O servidor deve retornar mensagens de erro apropriadas para cada
      situação (ex: ID inválido, método não permitido, compromisso não
      encontrado).
    - Os erros devem ser retornados com os status HTTP apropriados (ex:
      400 Bad Request, 404 Not Found, 405 Method Not Allowed, 500
      Internal Server Error).

**Instruções para o Desenvolvedor**

1. **Criação do Projeto**
    - Crie um novo projeto Go.
    - Implemente os arquivosmain.go,compromisso.goegerenciador-compromisso.go
      conforme as especificações fornecidas.
2. **Testes**
    - Execute o servidor e teste as funcionalidades de criar, listar, obter,


```
atualizar e excluir compromissos utilizando ferramentas comocurl
ou Postman.
```
- Certifique-se de que o sistema funciona conforme o esperado e corrija
  quaisquer erros encontrados.
3. **Documentação**
- Documente o código com comentários explicativos.
- Forneça exemplos de requisições e respostas para cada rota.

**Exemplos de Requisições e Respostas**

1. **Criar Compromisso**
    - Requisição:
      curl -X POST -H "Content-Type: application/json" -d
      '{"title": "Reunião de Equipe", "description": "Discussão sobre o projeto X",
      "date": "2023-10-01", "time": "10:00"}' [http://localhost:8080/compromissos](http://localhost:8080/compromissos)
    - Resposta:
      {
      "id": 1,
      "title": "Reunião de Equipe",
      "description": "Discussão sobre o projeto X",
      "date": "2023-10-01",
      "time": "10:00"
      }
2. **Listar Compromissos**
    - Requisição:
      curl -X GET [http://localhost:8080/compromissos](http://localhost:8080/compromissos)
    - Resposta:
      [
      {
      "id": 1,
      "title": "Reunião de Equipe",
      "description": "Discussão sobre o projeto X",
      "date": "2023-10-01",
      "time": "10:00"
      }
      ]
3. **Obter Compromisso por ID**
    - Requisição:
      curl -X GET [http://localhost:8080/compromissos/](http://localhost:8080/compromissos/)
    - Resposta:
      {
      "id": 1,
      "title": "Reunião de Equipe",
      "description": "Discussão sobre o projeto X",
      "date": "2023-10-01",
      "time": "10:00"


## }

4. **Atualizar Compromisso**
    - Requisição:
      curl -X PUT -H "Content-Type: application/json" -d
      '{"title": "Reunião de Equipe Atualizada",
      "description": "Nova discussão sobre o projeto X",
      "date": "2023-10-02", "time": "11:00"}' [http://localhost:8080/compromissos/](http://localhost:8080/compromissos/)
    - Resposta:
      {
      "id": 1,
      "title": "Reunião de Equipe Atualizada",
      "description": "Nova discussão sobre o projeto X",
      "date": "2023-10-02",
      "time": "11:00"
      }
5. **Excluir Compromisso**
    - Requisição:
      curl -X DELETE [http://localhost:8080/compromissos/](http://localhost:8080/compromissos/)
    - Resposta:
      **-** Status: 204 No Content

**Conclusão** Este documento fornece todas as especificações necessárias para que
o desenvolvedor possa implementar o sistema de gerenciamento de compromissos
conforme solicitado. Certifique-se de seguir todas as instruções e requisitos para
garantir a correta implementação do sistema.


