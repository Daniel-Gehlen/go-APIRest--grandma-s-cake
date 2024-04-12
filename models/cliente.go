package models

type Cliente struct {
    ID        int     `json:"id"`
    Nome      string  `json:"nome"`
    Endereco  string  `json:"endereco"`
    Telefone  string  `json:"telefone"`
    Email     string  `json:"email"`
    Bolos     []Bolo  `json:"bolos"` // Relação 1:N
}
