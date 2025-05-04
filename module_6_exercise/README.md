# 📅 Sistema de Gerenciamento de Compromissos (Go + Echo)


### ✅ Objetivo
Desenvolver um sistema de gerenciamento de compromissos que permita a criação, leitura, atualização e exclusão de compromissos. Cada compromisso deve conter um título, uma descrição, uma data e um horário. O sistema deve ser implementado em Go, utilizando Echo Framework, e seguir a estrutura e práticas de codificação fornecidas nos exemplos de código.

---

### 📋 Requisitos Funcionais
1. **Inicialização do Servidor** – O servidor deve ser inicializado na porta 8080.
2. **Gerenciador de Compromissos**
    - Deve ser inicializado com um mapa vazio de compromissos, um ID inicial de 1 e um canal para solicitações de compromissos.
    - Deve iniciar uma goroutine para processar solicitações de compromissos.
3. **Manipulação de Rotas**
    - Definir as seguintes rotas:
        - `/compromissos`: Para criar e listar compromissos.
        - `/compromissos/{id}`: Para obter, atualizar e excluir um compromisso específico.
4. **Criação de Compromissos**
    - A rota `/compromissos` deve aceitar requisições POST com JSON contendo: `title`, `description`, `date` e `time`.
    - Retornar o compromisso criado com status HTTP 201 Created.
5. **Listagem de Compromissos**
    - A rota `/compromissos` deve aceitar requisições GET.
    - Retornar lista de compromissos com status HTTP 200 OK.
6. **Obtenção por ID**
    - A rota `/compromissos/{id}` deve aceitar GET e retornar o compromisso ou HTTP 404.
7. **Atualização de Compromisso**
    - A rota `/compromissos/{id}` deve aceitar PUT com JSON atualizado e retornar HTTP 200 ou 404.
8. **Exclusão de Compromisso**
    - A rota `/compromissos/{id}` deve aceitar DELETE e retornar HTTP 204 ou 404.

---

### 📋  Requisitos Não Funcionais
1. **Estrutura do Projeto**
    - Arquivos: `main.go`, `compromisso.go`, `gerenciador-compromisso.go`.
2. **Práticas de Codificação**
    - Seguir os exemplos fornecidos. Não usar pacotes adicionais.
3. **Tratamento de Erros**
    - Mensagens e status HTTP apropriados: 400, 404, 405, 500.

---

## 🚀 Como Iniciar o Projeto

### Instale as dependências

```bash
go mod tidy
```

### Execute o servidor

```bash
go run main.go
```

O servidor será iniciado em `http://localhost:8080`.

---

## 🧪 Testes com cURL

### ✅ Criar um compromisso

```bash
curl -X POST http://localhost:8080/compromissos   -H "Content-Type: application/json"   -d '{
    "title": "Reunião de Equipe",
    "description": "Discussão sobre o projeto X",
    "date": "2023-10-01",
    "time": "10:00"
}'
```

### 📋 Listar compromissos

```bash
curl -X GET http://localhost:8080/compromissos
```

### 🔎 Obter compromisso por ID

```bash
curl -X GET http://localhost:8080/compromissos/1
```

### ✏️ Atualizar compromisso

```bash
curl -X PUT http://localhost:8080/compromissos/1   -H "Content-Type: application/json"   -d '{
    "title": "Reunião Atualizada",
    "description": "Nova discussão",
    "date": "2023-10-02",
    "time": "11:00"
}'
```

### ❌ Excluir compromisso

```bash
curl -X DELETE http://localhost:8080/compromissos/1
```

---

## 📌 Rotas da API

| Método | Rota                    | Descrição                         |
|--------|-------------------------|-----------------------------------|
| POST   | /compromissos           | Criar novo compromisso            |
| GET    | /compromissos           | Listar todos os compromissos      |
| GET    | /compromissos/:id       | Obter compromisso específico      |
| PUT    | /compromissos/:id       | Atualizar compromisso existente   |
| DELETE | /compromissos/:id       | Excluir compromisso existente     |

---

### Conclusão
Este documento fornece todas as especificações necessárias para que o desenvolvedor possa implementar o sistema de gerenciamento de compromissos conforme solicitado.