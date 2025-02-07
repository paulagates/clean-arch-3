# Clean Architecture Challenge - Order Management System

Este projeto é um desafio de implementação de uma aplicação seguindo os princípios da Clean Architecture. A aplicação consiste em um sistema de gerenciamento de pedidos (orders) que expõe três tipos de interfaces para listar pedidos: REST, gRPC e GraphQL. Além disso, o projeto utiliza Docker para facilitar a configuração do ambiente de desenvolvimento.

## Requisitos

- **Endpoint REST**: `GET /order`
- **Service ListOrders**: Implementado com gRPC
- **Query ListOrders**: Implementado com GraphQL
- **Migrações**: Criação das migrações necessárias para o banco de dados
- **Arquivo `api.http`**: Contendo as requisições para criar e listar pedidos
- **Docker**: Utilização de `Dockerfile` e `docker-compose.yaml` para subir o banco de dados MySQL e o RabbitMQ

## Tecnologias Utilizadas

- **Go**
- **MySQL**
- **RabbitMQ**
- **gRPC**
- **GraphQL**
- **Docker**

---

## Passos para Execução

### Passo 1: Subir o Ambiente com Docker

Para iniciar o ambiente de desenvolvimento, execute o seguinte comando:

docker-compose up

### Passo 2: Executar as Migrações

Após subir o banco de dados, execute as migrações para criar as tabelas necessárias:
``` 
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" up
``` 
### Passo 3: Cadastrar uma Order

Utilize o arquivo api.http para realizar a requisição de criação de uma order. Exemplo:
``` 
POST http://localhost:8080/order
Content-Type: application/json

{
  "price": 100.50,
  "tax": 10.05
}
``` 
### Passo 4: Listar Orders

Você pode listar as orders utilizando uma das três interfaces disponíveis:

## 1. REST (HTTP)
``` 
GET http://localhost:8080/order
``` 
## 2. GraphQL

No playground do GraphQL, utilize a seguinte query:
``` 
query queryOrders {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}
``` 
## 3. gRPC

# Para testar o serviço gRPC, utilize o Evans CLI:
``` 
evans -r repl
``` 
# Dentro do Evans, siga os passos:

- Selecionar o serviço ListOrders:
``` 
service ListOrders
``` 

- Chamar o método ListOrders:
``` 
call ListOrders
``` 

### Portas dos Serviços

- **REST API:** 8000

- **gRPC:** 50051

- **GraphQL Playground:** 8080 (ou outra porta configurada)

- **MySQL:** 3306

- **RabbitMQ:** 5672 (AMQP) e 15672 (Management UI)

