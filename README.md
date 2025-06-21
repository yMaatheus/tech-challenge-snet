# Tech Challenge SNET API

API para gerenciamento de estabelecimentos e lojas.

---

## 🏗️ Como rodar o projeto

### Localmente

1. Clone o repositório:
    ```
    git clone https://github.com/yMaatheus/tech-challenge-snet.git
    cd tech-challenge-snet
    ```

2. Configure os arquivos dentro de `web/` e `server/` `.env` (pode copiar de `.env.example`):

    ```
    cp .env.example .env
    ```

    Edite as variáveis conforme seu ambiente

3. Instale as dependências:

    ```
    go mod tidy
    ```

4. Rode as migrações do banco:

    ```
    make migrate
    ```

5. Inicie o servidor local:
    ```
    make dev

    ou

    make run
    ```

### Via Docker

Siga os passos 1 ao 4 antes de seguir via docker.

1. Suba os containers:
    ```
    docker-compose up --build
    ```

2. O backend estará em: http://localhost:8080 e o frontend em: http://localhost:3000

---

## 🛠️ Scripts úteis do Makefile

- make dev         # Sobe o servidor em modo desenvolvimento (hot reload)
- make build       # Compila o projeto
- make run         # Roda o binário compilado
- make test        # Executa todos os testes
- make coverage    # Gera o relatório de cobertura dos testes
- make migrate     # Executa as migrations do banco de dados
- make docker      # Builda e sobe tudo com Docker

---

## 🧪 Como rodar os testes

    make test

  Para relatório de cobertura:

    make coverage

---

## 📑 Documentação Swagger

- Com o servidor rodando acesse: http://localhost:8080/docs

---

## 🚀 Principais Endpoints

### Estabelecimentos

- POST `/establishments`
    Cria um novo estabelecimento.

    Exemplo de body:
    ```
    {
        "number": "E001",
        "name": "Restaurante XPTO",
        "corporate_name": "Restaurante XPTO LTDA",
        "address": "Rua Exemplo",
        "address_number": "99",
        "city": "São Paulo",
        "state": "SP",
        "zip_code": "01001000"
    }
    ```

- GET `/establishments`
    Lista estabelecimentos e o campo storesTotal.

- GET `/establishments/{id}`
    Detalhes do estabelecimento e suas lojas.

- PUT `/establishments/{id}`
    Atualiza um estabelecimento.

- DELETE `/establishments/{id}`
    Remove um estabelecimento (apenas se não houver lojas).

### Lojas

- POST   /stores
- GET    /stores
- GET    /stores/{id}
- PUT    /stores/{id}
- DELETE /stores/{id}

---

## ⚠️ Exemplos de Resposta de Erro

- Validação:
    {
        "validation_error": {
            "Campo": "motivo"
        }
    }

- Parâmetro inválido:
    {
        "error": "Invalid establishment ID. Must be a positive integer."
    }

---
