# Tech Challenge Snet

API para gerenciamento de estabelecimentos e lojas.

## Deploy

- **URL de Produção:**
  - Frontend: [https://snet.ymaatheus.dev/](https://snet.ymaatheus.dev/)
  - Backend: [https://snet-api.fly.dev/](https://snet-api.fly.dev/)
  
  > Documentação Swagger: https://snet-api.fly.dev/docs

  ## Atenção! A Autenticação no FrontEnd é apenas uma simulação!
  Digite qualquer email e senha e você conseguirá acessar o dashboard da aplicação.

---

## 💡 Possíveis Melhorias
- Separação de DTOs (requests/responses) dos Models de domínio.
- Mensagens de validação customizadas e validação condicional nos endpoints.
- Cobertura de testes para cenários de erro e casos de borda.
- Paginação, busca e filtros nos endpoints de listagem.
- Documentação Swagger com exemplos de erros e payloads.
- Implementação de autenticação/autorização (ex: JWT).
- Uso de migrations versionadas (golang-migrate).
- Adição de métricas para monitoramento e tracing (OpenTelemetry).
- Melhor organização das camadas (UseCases).
- Validação nos formulários do frontend para garantir melhor UX.

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

# 🏗️ Estrutura de Pastas do Backend

```

server/
  cmd/
    main.go               # Ponto de entrada da aplicação. Inicializa Echo, middlewares, rotas e configs principais.
  config/
    config.go             # Gerenciamento de variáveis de ambiente e configuração de banco de dados.
  database/
    migration.sql         # Scripts de migração e estruturação do banco de dados.
    reset.sql             # Script para resetar o banco em ambiente de desenvolvimento/teste.
  docs/
    ...                   # Arquivos do Swagger/OpenAPI. Documentação da API (acessível em /docs).
  handler/
    ...                   # Handlers: camada responsável por processar as requisições HTTP, validar dados e retornar respostas.
  model/
    ...                   # Models/Entidades: Definições das structs usadas em todo o sistema (ex: Store, Establishment).
  repository/
    ...                   # Repositórios: Camada de acesso ao banco de dados, SQL queries e CRUD.
  service/
    ...                   # Serviços (business logic): Orquestração das regras de negócio do sistema.
  testutil/
    ...                   # Utilitários para facilitar a execução de testes, como helpers para banco de dados e mocks.
  util/
    ...                   # Funções auxiliares gerais: helpers de validação, formatação de erros, etc.
  go.mod
  go.sum

```

# 📑 Frontend NuxtJs 3 - Estrutura de Pastas

```
src/
├── assets/                   # Arquivos estáticos (imagens, css, fontes)
│   └── css/
│       └── tailwind.css
├── components/               
│   ├── app/                  # Componentes com regra de negócio (ex: formulários de domínio)
│   └── ui/                   # Design system, componentes visuais reutilizáveis
│       ├── button/
│       ├── ...
├── layouts/                  # Layouts globais
├── lib/                      # Funções utilitárias e helpers
│   └── utils.ts
├── middleware/               # Middlewares (ex: autenticação)
│   └── auth.ts
├── pages/                    # Páginas do projeto, estrutura reflete as rotas
│   ├── index.vue
│   └── dashboard/
│       └── ...
├── public/                   # Arquivos públicos acessíveis por URL (favicon, robots.txt)
├── services/                 # Serviços para chamadas de API e lógica de integração
├── types/                    # Tipos e interfaces TypeScript do domínio
└── ...

```
