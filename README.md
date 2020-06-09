# Mercado Livre Crawler
Esse projeto consiste em uma api rest usada para extrair dados da página html do resultado da pesquisa de produtos do site [https://www.mercadolivre.com.br/](https://www.mercadolivre.com.br/). Você pode pesquisar qualquer produto, os resultados da pesquisa serão analisados e as informações como nome do produto, link da página do produto, preço, loja que está vendendo e localização serão extraídas do html e transformadas em JSON, através de um endopint rest.

## Tecnologias
- Go v1.13.8
- Git

## API
URI base: http://localhost:3000/api

|      Endpoints    |       Método HTTP       |        Corpo da Requisição          |   Parâmetros Obrigatórios  |  Valores Padrão  |
|-------------------|-------------------------|-------------------------------------|----------------------------|------------------| 
|    /v1/products   |         POST            | `{ "search": "Casa", "limit": 20 }` |           `search`         |    `limit = 5`   |


## Como executar

> Comandos executados em sistema baseado em Unix.
> Necessário instalar as ferramentas citadas no item de tecnologias.

a) Clonar o repositório
- `git clone https://github.com/rof20004/mercadolivre-crawler-go.git`

b) Instalar as dependências
- `go mod tidy`

c) Subir a aplicação
- `go run *.go`

## Exemplo de utilização
`curl -i -X POST http://localhost:3000/api/v1/products -d '{ "search": "playstation4" }' -H "content-type: application/json"`

**Obs.: é obrigatório informar o content-type na requisição**