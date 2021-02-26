# bankingApp

## Imports used
- encoding/json         ()
- fmt                   ()
- net/http              ()
- gorilla mux           ()
- database/sql          ()
- go-sql-driver/mysql   ()
- log                   (logs)
- go.uber.org/zap       (logger)

## Routes
- GET  - /customers
- GET  - /customers/{id}
- POST - /customers/{customer_id}/account               { "account_type" : ['saving','checking'] }
- POST - /customers/{customer_id}/account/{account_id}  




# What I learned
* Apply SOLID design principles in Go
* Hexagonal Architecture design and its implementation
* Implement authentication and authorization using JWT tokens
* Apply RBAC Authorization to APIs
* Understand how dependency injection works in Go
* Understand and implement the structured logging
* Build microservices API in Go
* Code Refactoring in Go
* Decoupling the domain objects and DTOs
* Working with small steps
* Take the informed decision on choosing various libraries
* Understand the role of multiplexer in HTTP web server
* Encoding structs to JSON or XML
* Understand the routing capabilities of gorilla/mux
* Design your own error library
* State based unit testing
* Unit testing using the mocks

# Hexagonal architecture
![alt text](https://i.ibb.co/RCKw5Gg/Hex-arq-repo-data-adapter.png)


