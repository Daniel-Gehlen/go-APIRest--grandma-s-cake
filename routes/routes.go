package routes

import (
    "net/http"

    "api-bolos-vovo/handlers"

    "github.com/gorilla/mux"
)

// NewRouter: Cria um novo roteador Gorilla Mux
func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Rotas para Bolos
    router.HandleFunc("/bolos", handlers.GetBolos).Methods(http.MethodGet)
    router.HandleFunc("/bolos", handlers.CreateBolo).Methods(http.MethodPost)
    // ... (adicionar rotas para GetBoloByID, UpdateBolo, DeleteBolo)

    // Rotas para Clientes
    router.HandleFunc("/clientes", handlers.GetClientes).Methods(http.MethodGet)
    router.HandleFunc("/clientes", handlers.CreateCliente).Methods(http.MethodPost)
    // ... (adicionar rotas para GetClienteByID, UpdateCliente, DeleteCliente)

    return router
}
