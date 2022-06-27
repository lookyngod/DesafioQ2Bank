##  Desafio/Objetivo:
Temos 2 tipos de usuários, os comuns e lojistas, ambos têm carteira com dinheiro e
realizam transferências entre eles. Vamos nos atentar somente ao fluxo de
transferência entre dois usuários.

A aplicação API tem como responsabilidade fornecer uma tabela com CPF e CNPJ quando inseridos.

## Stack

- Golang
- Docker
- Docker Compose
 
 PS: Para logs direto do BackEnd, recomendo usar Postman com os endpoints inseridos no Router.go


## Instalação

A aplicação necessite de um ambiente com [Golang](https://go.dev/doc/install) 1.17+ para rodar
e do [Docker Compose](https://docs.docker.com/compose/install/)


Instale as dependências e para rodar a aplicação use o passo-a-passo abaixo:

### Passo 1:
Suba os containers rodando o comando abaixo na pasta raiz do projeto:
```sh
docker-compose up
```

### Passo 2:
Após a inicialização do docker-compose, entre na pasta '/api' e em um novo terminal, rode o comando:
```sh
go run router.go
```

### Passo 3:

- *Prontinho, sua aplicação está no ar!*

## Como testar a aplicação?
- Por padrão sua aplicação front irá rodar na porta 5000 ```http://localhost:5000/api/{endpoints}```
- A aplicação contém um prefixo rota '/api' para teste via Postman com arquivo JSON ```http://localhost:5000/api```


# Métodos
Requisições para a API devem seguir os padrões:
| Método | Descrição |
|---|---|
| `GET` | Retorna informações de usuários na resposta |
| `POST` | Utilizado para inserir uma transação no banco de dados|

## Notes
