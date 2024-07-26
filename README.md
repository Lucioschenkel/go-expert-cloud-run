# Introdução

Este repositório contém o código referente ao laboratório "Deploy com Google Cloud Run" da pós graduação Go Expert, da Full Cycle.

# Executando localmente

Para rodar este projeto localmente, existem duas opções:

1. Usando `go` diretamente (requer versão 1.22.2)
2. Usando `docker-compose` (recomendado)

## Executando com `go`

Para executar o programa, primeiro é preciso instalar as dependências. Para isso, execute:

```bash
go mod tidy
```

Após instalar as dependências, o projeto pode ser executado normalmente, usando o comando:

`go run cmd/server/main.go`

Pronto! O servidor estará disponível na porta 8080. Você pode testar o projeto fazendo uma requisição `GET` para a rota `/api/weather/{cep}`, ex:

```bash
curl http://localhost:8080/api/weather/70165900
```

:::info

Nota: é preciso configurar a variável de ambiente `API_KEY` para que o programa seja executado corretamente. Você pode gerar um chave de API no site https://www.weatherapi.com.

:::

## Executando com Docker compose

Para executar o programa usando Docker, siga os seguintes passos:

1. Crie um arquivo .env:

```bash
cp .env.example .env
```

2. Substitua o valor da variável `API_KEY` no arquivo `.env` por sua própria chave de API.
3. Execcute o comando para subir a aplicação

```bash
docker compose up -d
```

Pronto! O servidor estará disponível na porta 8080. Você pode testar o projeto fazendo uma requisição `GET` para a rota `/api/weather/{cep}`, ex:

```bash
curl http://localhost:8080/api/weather/70165900
```

# Testes automatizados

Como esse programa têm dependências em serviços externos, foram incluídos testes de integração para o `usecase` da aplicação (busca da temperatura atual para um determinado CEP). Por esse motivo, se faz necessário o uso da chave de API do serviço `weatherapi.com` para testar o projeto.

Gere sua chave de API diretamente no site do serviço, e configure uma variável de ambiente chamada `API_KEY` cujo valor é a sua própria chave de API do serviço `weatherapi.com`, e então execute o comando:

```bash
go test ./...
```

# Acessando a aplicação hospedada no Google Cloud Run

Esse projeto está hospedado no Google Cloud Run na seguinte URL: https://go-expert-cloud-run-3xubhsjt5q-uc.a.run.app

Para testar a aplicação, utilize a seguinte rota:

`https://go-expert-cloud-run-3xubhsjt5q-uc.a.run.app/api/weather/{CEP}`

Na URL acima, substitua `{CEP}` pelo CEP desejado para a consulta da temperatura.

Exemplo:

```bash
curl https://go-expert-cloud-run-3xubhsjt5q-uc.a.run.app/api/weather/70165900
```
