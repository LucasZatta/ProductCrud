# ProductCrud

## Project Structure:

```
├── cmd
│   └──main  
|      └──main.go      \\ Server startup
├── internal
│   └──database        \\ Database layer, firs implemented as a service
|      └──database.go
│   └──products         \\ Everything related to the product entity  
|      └──controller.go \\ Controller layer to handle product related requests
|      └──model.go      \\ Product entity model
|      └──repository.go \\ Repository layer to handle product related db operations
│   └──server
|      └──routes.go     \\ Where routes and handlers are setup
|      └──server.go     \\ Server config and setups
├── docker-compose.yml  \\ docker-compose instructions for mySQL db container
├── go.mod              \\ Go dependencies file
├── go.sum
├── init.sql            \\ Init sql file to create table on container creation
├── makefile            
```

## Considerations

### Structure
For this project I opted to test out new project layouts and strucutures. I had in min setting up a database layer as a service, but that proved to be a bit of an overkill and it ended up making things more difficult than they needed to be. I ended up with something ressembling a Controller, Service, Repository pattern(withou the Service layer and some pattern conventions), with the database instance being passed arround throughout the code wherever needed using the server struct + database interface. 

### Echo Framework
Was not sure how to handle erros following what was seen on the docs. Followed what seemed like the convention for Echo but was not quite happy with the way the error handling ended up in this project. Overall had a bit of fun testing echo out.

## Rest API
### GET    /products/:id
This endpoint uses the id parameter to find and return its corresponding product in the database.
### POST   /products
Using form-values, this endpoint creates a new product and returns its id.
### DELETE /products/:id
This endpoint uses the id parameter to delete its corresponding product in the database.
### PUTH   /products/:id
Using both form-values and the id parameter, this enpoint updates the correspondy entry in the database with new values.

## Running the project
First use the command make docker-run to start the db container, it will use the init.sql to create the table needed for testing.

Then use make run to run the project.

👍 All api tests were done using postman or insomnia 👍
