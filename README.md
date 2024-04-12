# API REST for Grandma's Cakes with Gorilla Mux

**Objective:**

The aim was to create a comprehensive REST API for managing cakes and clients within a virtual bakery. The API needed to support various operations:

**Cakes:**
- Listing all cakes.
- Creating a new cake.
- Searching for a cake by ID.
- Updating a cake.
- Deleting a cake.

**Clients:**
- Listing all clients.
- Creating a new client.
- Searching for a client by ID.
- Updating a client.
- Deleting a client.

**Implementation:**

**Project Structure:**
The project was organized into folders for models (cakes and clients), handlers (data manipulation), and routes (route definitions).

**Models:**
The Cake and Client structures were defined with their respective attributes.

**Data Manipulation:**
Handler functions were created for each CRUD operation (Create, Read, Update, Delete) for cakes and clients.
These functions utilized an in-memory slice (cakes and clients) to store data.

**Routes:**
Routes were defined in the routes file, mapping HTTP routes for each CRUD operation using Gorilla Mux.
HTTP verbs (GET, POST, PUT, DELETE) were mapped to the appropriate handler functions.

**Server:**
The main file configured and initiated the HTTP server, specifying the port and utilizing the



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
