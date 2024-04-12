package handlers

import (
    "encoding/json"
    "net/http"

    "api-bolos-vovo/models"
)

var bolos []models.Bolo

// GetBolos: Retorna todos os bolos
func GetBolos(w http.ResponseWriter, r *http.Request) {
    jsonData, err := json.Marshal(bolos)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// CreateBolo: Cria um novo bolo
func CreateBolo(w http.ResponseWriter, r *http.Request) {
    var novoBolo models.Bolo
    err := json.NewDecoder(r.Body).Decode(&novoBolo)
    if err != nil {
        http.Error(w, "Erro ao ler o JSON do bolo", http.StatusBadRequest)
        return
    }

    bolos = append(bolos, novoBolo)

    jsonData, err := json.Marshal(novoBolo)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write(jsonData)
}

// ... (implementar outros métodos para GetBoloByID, UpdateBolo, DeleteBolo)
