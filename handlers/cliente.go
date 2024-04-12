package handlers

import (
    "encoding/json"
    "net/http"

    "api-bolos-vovo/models"
)

var clientes []models.Cliente

// GetClientes: Retorna todos os clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {
    jsonData, err := json.Marshal(clientes)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// CreateCliente: Cria um novo cliente
func CreateCliente(w http.ResponseWriter, r *http.Request) {
    var novoCliente models.Cliente
    err := json.NewDecoder(r.Body).Decode(&novoCliente)
    if err != nil {
        http.Error(w, "Erro ao ler o JSON do cliente", http.StatusBadRequest)
        return
    }

    clientes = append(clientes, novoCliente)

    jsonData, err := json.Marshal(novoCliente)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write(jsonData)
}

// ... (
