## API REST for Grandma's Cakes with Gorilla Mux

Note: This is a simplified example of a REST API. A complete implementation would involve more features, data validation, security, and a database.

1. Installation of Gorilla Mux:

```bash
go get -u github.com/gorilla/mux
```

## Project Structure:

```
api-bolos-vovo/
├── main.go
├── models/
│   ├── bolo.go
│   └── cliente.go
├── handlers/
│   ├── bolos.go
│   └── clientes.go
└── routes/
    └── routes.go
```

## main.go for Grandma's Cakes REST API
This file sets up the server and defines the routes of the API:

## Explanation:

Package Imports:

fmt: For output formatting.
log: For logging messages to the console.
net/http: For HTTP functionalities.
api-bolos-vovo/routes: Imports the routes package which defines the API routes.
Server Configuration:

router := routes.NewRouter(): Creates a new router using the NewRouter function from the routes package.
port := ":8080": Defines the server port (can be changed).
Server Start:

fmt.Printf("... running on port %s\n", port): Prints a message indicating the port the server is running on.
log.Fatal(http.ListenAndServe(port, router)): Starts the HTTP server listening on the specified port and assigns the router to handle requests.
Note:

The implementation of the routes package contains the logic to map HTTP verbs (GET, POST, PUT, DELETE) to the handlers defined in handlers/bolos.go and handlers/clientes.go.
