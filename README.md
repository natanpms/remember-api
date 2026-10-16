# Remember API

API REST desenvolvida para gerenciar contéudo de uma aplicação web de rede neurais com foco em criar, gerenciar e cultivar memórias e ideias.

## Tecnologias

- [Go](https://go.dev/) + [Gin](https://gin-gonic.com/) — servidor HTTP
- [PostgreSQL 16](https://www.postgresql.org/) — banco de dados
- [Docker](https://www.docker.com/) + Docker Compose — ambiente containerizado
- [Air](https://github.com/air-verse/air) + Compose Watch — hot reload em desenvolvimento
- [Swag](https://github.com/swaggo/swag) — geração da documentação Swagger

## Pré-requisitos

- Ambiente Docker

## Configuração

1. Clone o repositório:

```bash
git clone <url-do-repositorio>
cd remember_api
```

2. Crie o arquivo `.env` a partir do exemplo:

```bash
cp .env.example .env
```

## Como rodar

### Desenvolvimento (com hot reload) — recomendado

```powershell
docker compose -f docker-compose.yml -f docker-compose.dev.yml up --watch
```

> A API leva ~10-15 segundos para ficar pronta (o Air compila o projeto na subida). Para acompanhar: `docker logs remember-api -f`

### Produção (binário compilado)

```powershell
docker compose up -d --build
```

## Documentação (Swagger)

Com a API rodando, acesse a documentação interativa no navegador:

-  http://localhost:8000/api/v1/swagger/index.html

Por lá é possível visualizar todos os endpoints.

### Regenerando a documentação

Sempre que adicionar ou alterar anotações Swag nos handlers, regenere os arquivos em `docs/`:

```powershell
swag init -g cmd/main.go
```

