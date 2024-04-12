package models

type Bolo struct {
    ID        int     `json:"id"`
    Sabor     string  `json:"sabor"`
    Tamanho   string  `json:"tamanho"`
    Preco     float64 `json:"preco"`
    ClienteID int     `json:"cliente_id"`
}
