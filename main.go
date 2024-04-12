package main

import (
    "fmt"
    "log"
    "net/http"

    "api-bolos-vovo/routes"
)

func main() {
    // Inicializa rotas
    router := routes.NewRouter()

    // Define porta do servidor
    port := ":8080"

    fmt.Printf("Servidor da API Bolos da Vovó rodando na porta %s\n", port)
    log.Fatal(http.ListenAndServe(port, router))
}
