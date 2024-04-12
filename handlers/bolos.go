package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

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

    novoBolo.ID = len(bolos) + 1 // Gera ID automaticamente
    bolos = append(bolos, novoBolo)

    jsonData, err := json.Marshal(novoBolo)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write(jsonData)
}

// GetBoloByID: Retorna um bolo por ID
func GetBoloByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    bolo, ok := findBoloByID(id)
    if !ok {
        http.Error(w, "Bolo não encontrado", http.StatusNotFound)
        return
    }

    jsonData, err := json.Marshal(bolo)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// UpdateBolo: Atualiza um bolo
func UpdateBolo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var boloAtualizado models.Bolo
    err = json.NewDecoder(r.Body).Decode(&novoBolo)
    if err != nil {
        http.Error(w, "Erro ao ler o JSON do bolo", http.StatusBadRequest)
        return
    }

    boloAtualizado.ID = id // Mantém o ID original

    updated, ok := updateBolo(id, boloAtualizado)
    if !ok {
        http.Error(w, "Erro ao atualizar bolo", http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(updated)
    if err != nil {
        http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// DeleteBolo: Deleta um bolo
func DeleteBolo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    deleted, ok := deleteBolo(id)
    if !ok {
        http.Error(w, "Bolo não encontrado", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// Funções auxiliares para manipulação interna do slice de bolos
func findBoloByID(id int) (models.Bolo, bool) {
    for _, bolo := range bolos {
        if bolo.ID == id {
            return bolo, true
        }
    }
    return models.
